package controller

import (
	"fmt"
	"messenger/internal/models"
	"messenger/internal/utils"
	"net/http"

	"github.com/rs/zerolog/log"
)

func (hm *HandlersManager) GetChatsHandler() map[string]string {
	resp := make(map[string]string)

	response, err := hm.client.R().
		SetHeader("Authorization", utils.GetToken()).
		SetResult(&resp).
		Get(fmt.Sprintf("https://%s/chats", hm.addr))

	if err != nil {
		log.Error().Err(err).Msg("cant request getting chats")
		return nil
	}

	if response.StatusCode() != http.StatusOK {
		log.Error().Msg(string(response.Body()))
	}

	return resp
}

func (hm *HandlersManager) ChatHistoryHandler(chatId string) []models.ShortMessage {
	var resp []models.ShortMessage

	response, err := hm.client.R().
		SetHeader("Authorization", utils.GetToken()).
		SetResult(&resp).
		Get(fmt.Sprintf("https://%s/chats/%s/history", hm.addr, chatId))

	if err != nil {
		log.Error().Err(err).Msg("cant request getting chats")
		return nil
	}

	if response.StatusCode() != http.StatusOK {
		log.Error().Msg(string(response.Body()))
	}

	return resp
}
