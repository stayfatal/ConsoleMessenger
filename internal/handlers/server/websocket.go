package handlers

import (
	"messenger/internal/database"
	"messenger/internal/parsers"
	websocket "messenger/internal/websocket/server"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (hm *handlersManager) NewChatHandler(c *gin.Context) {
	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	val, ok := c.Get("id")
	if !ok {
		log.Error().Err(err).Msg("cant get id from token")
		return
	}

	id, err := parsers.FromAnyToUUID(val)
	if err != nil {
		log.Error().Err(err).Msg("cant convert uuid")
		return
	}

	websocket.AddConn(id, conn)

	recipientName := c.GetHeader("Recipient")

	recipient, err := hm.dm.GetUserByName(recipientName)
	if err != nil {
		log.Error().Err(err).Msg("cant get user by his name")
		return
	}

	chatId := uuid.New()
	err = hm.dm.CreateChat(database.Chat{
		Id:      chatId,
		User1Id: id,
		User2Id: recipient.Id,
	})

	if err != nil {
		log.Error().Err(err).Msg("cant create chat")
		return
	}

	websocket.StartChat(id, recipient.Id)
}

func (hm *handlersManager) JoinChatHandler(c *gin.Context) {
	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	id, err := uuid.Parse(c.GetString("id"))
	if err != nil {
		log.Error().Err(err).Msg("cant parse uuid from context")
		return
	}

	websocket.AddConn(id, conn)

	chatId, err := uuid.Parse(c.Params.ByName("chat_id"))
	if err != nil {
		log.Error().Err(err).Msg("cant parse uuid")
		return
	}

	chat, err := hm.dm.GetChatById(chatId)
	if err != nil {
		log.Error().Err(err).Msg("cant get chat by id")
		return
	}

	websocket.StartChat(chat.User1Id, chat.User2Id)
}
