package handlers

import (
	"fmt"
	"messenger/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (hm *handlersManager) NewChatHandler(c *gin.Context) {
	hm.wm.Upgrade(c)

	senderId := c.GetInt("id")

	sender, err := hm.dm.GetUserById(senderId)
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling GetUserById")).Msg("")
		return
	}

	recipientUsername := c.GetHeader("Recipient")

	recipient, err := hm.dm.GetUserByName(recipientUsername)
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling GetUserByName")).Msg("")
		return
	}

	chatId, err := hm.dm.CreateChat(models.Chat{
		Name: fmt.Sprintf("%s and %s", sender.Username, recipient.Username),
	})
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling CreateChat")).Msg("")
		return
	}

	err = hm.dm.AddChatMember(models.ChatMember{
		ChatId: chatId,
		UserId: sender.Id,
	})
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling AddChatMember")).Msg("")
		return
	}

	err = hm.dm.AddChatMember(models.ChatMember{
		ChatId: chatId,
		UserId: recipient.Id,
	})
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling AddChatMember")).Msg("")
		return
	}

	hm.wm.JoinChat(senderId, chatId)
}

func (hm *handlersManager) JoinChatHandler(c *gin.Context) {
	err := hm.wm.Upgrade(c)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		log.Error().Stack().Err(errors.Wrap(err, "upgrading to websocket")).Msg("")
		return
	}

	chatId, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "parsing string to int")).Msg("")
		return
	}

	senderId := c.GetInt("id")

	hm.wm.JoinChat(senderId, chatId)
}
