package handlers

import (
	"net/http"

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
