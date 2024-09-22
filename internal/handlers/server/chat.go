package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (hm *handlersManager) GetChatsHandler(c *gin.Context) {
	id := c.GetInt("id")

	chats, err := hm.dm.GetAllUserChats(id)
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling GetAllUserChats")).Msg("")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, chats)
}

func (hm *handlersManager) GetLastChatMessagesHandler(c *gin.Context) {
	chatId, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "parsing string to int")).Msg("")
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	messages, err := hm.dm.GetLastChatMessages(chatId)
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling GetLastChatMessages")).Msg("")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, messages)
}
