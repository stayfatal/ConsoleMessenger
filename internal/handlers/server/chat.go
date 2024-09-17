package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (hm *handlersManager) GetChatsHandler(c *gin.Context) {
	id := c.GetInt("id")

	chats, err := hm.dm.GetAllUserChats(id)
	if err != nil {
		log.Error().Err(err).Msg("cant get user's chats")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, chats)
}

func (hm *handlersManager) GetLastChatMessagesHandler(c *gin.Context) {
	chatId, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		log.Error().Err(err).Msg("cant parse string to int")
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	messages, err := hm.dm.GetLastChatMessages(chatId)
	if err != nil {
		log.Error().Err(err).Msg("cant get messages from db")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, messages)
}
