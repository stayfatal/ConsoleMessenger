package handlers

import (
	"fmt"
	"messenger/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (hm *handlersManager) NewChatHandler(c *gin.Context) {
	hm.wm.Upgrade(c)

	senderId := c.GetInt("id")

	sender, err := hm.dm.GetUserById(senderId)
	if err != nil {
		log.Error().Err(err).Msg("cant get user by his id")
		return
	}

	recipientUsername := c.GetHeader("Recipient")

	recipient, err := hm.dm.GetUserByName(recipientUsername)
	if err != nil {
		log.Error().Err(err).Msg("cant get user by his name")
		return
	}

	chatId, err := hm.dm.CreateChat(database.Chat{
		Name: fmt.Sprintf("%s and %s", sender.Username, recipient.Username),
	})
	if err != nil {
		log.Error().Err(err).Msg("cant create chat")
		return
	}

	err = hm.dm.AddChatMember(database.ChatMember{
		ChatId: chatId,
		UserId: sender.Id,
	})
	if err != nil {
		log.Error().Err(err).Msg("cant add member to chat")
		return
	}

	err = hm.dm.AddChatMember(database.ChatMember{
		ChatId: chatId,
		UserId: recipient.Id,
	})
	if err != nil {
		log.Error().Err(err).Msg("cant add member to chat")
		return
	}

	hm.wm.StartChat(chatId)
}

func (hm *handlersManager) JoinChatHandler(c *gin.Context) {
	hm.wm.Upgrade(c)

	chatId, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		log.Error().Err(err).Msg("cant parse chat_id")
		return
	}

	hm.wm.StartChat(chatId)
}
