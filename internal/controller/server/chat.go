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

func (cr *Controller) GetChats(c *gin.Context) {
	id := c.GetInt("id")

	chats, err := cr.repo.GetAllUserChats(id)
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling GetAllUserChats")).Msg("")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, chats)
}

func (cr *Controller) GetLastChatMessages(c *gin.Context) {
	chatId, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "parsing string to int")).Msg("")
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	messages, err := cr.repo.GetLastChatMessages(chatId)
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling GetLastChatMessages")).Msg("")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, messages)
}

func (cr *Controller) NewChat(c *gin.Context) {
	cr.cs.Upgrade(c)

	senderId := c.GetInt("id")

	sender, err := cr.repo.GetUserById(senderId)
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling GetUserById")).Msg("")
		return
	}

	recipientUsername := c.GetHeader("Recipient")

	recipient, err := cr.repo.GetUserByName(recipientUsername)
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling GetUserByName")).Msg("")
		return
	}

	chatId, err := cr.repo.CreateChat(models.Chat{
		Name: fmt.Sprintf("%s and %s", sender.Username, recipient.Username),
	})
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling CreateChat")).Msg("")
		return
	}

	err = cr.repo.AddChatMember(models.ChatMember{
		ChatId: chatId,
		UserId: sender.Id,
	})
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling AddChatMember")).Msg("")
		return
	}

	err = cr.repo.AddChatMember(models.ChatMember{
		ChatId: chatId,
		UserId: recipient.Id,
	})
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling AddChatMember")).Msg("")
		return
	}

	cr.cs.JoinChat(senderId, chatId)
}

func (cr *Controller) JoinChat(c *gin.Context) {
	err := cr.cs.Upgrade(c)
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

	cr.cs.JoinChat(senderId, chatId)
}
