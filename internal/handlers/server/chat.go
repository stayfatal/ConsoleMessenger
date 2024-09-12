package handlers

import (
	"messenger/internal/parsers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (hm *handlersManager) GetChatsHandler(c *gin.Context) {
	val, _ := c.Get("id")

	id, err := parsers.FromAnyToUUID(val)
	if err != nil {
		log.Error().Err(err).Msg("cant parse uuid")
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	chats, err := hm.dm.GetAllUserChats(id)
	if err != nil {
		log.Error().Err(err).Msg("cant parse uuid")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, chats)
}
