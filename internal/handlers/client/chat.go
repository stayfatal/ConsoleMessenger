package handlers

import (
	"fmt"
	"messenger/internal/env"

	"github.com/rs/zerolog/log"
)

type ShortMessage struct {
	SenderUsername string `json:"sender_username"`
	Message        string `json:"message"`
}

func (hm *HandlersManager) GetChatsHandler() map[string]string {
	resp := make(map[string]string)

	_, err := hm.client.R().
		SetHeader("Authorization", env.GetToken()).
		SetResult(&resp).
		Get(fmt.Sprintf("https://%s/chats", hm.addr))

	if err != nil {
		log.Error().Err(err).Msg("cant request getting chats")
		return nil
	}

	return resp
}

func (hm *HandlersManager) ChatHistoryHandler(chatId string) []ShortMessage {
	var resp []ShortMessage

	_, err := hm.client.R().
		SetHeader("Authorization", env.GetToken()).
		SetResult(&resp).
		Get(fmt.Sprintf("https://%s/chats/%s/history", hm.addr, chatId))

	if err != nil {
		log.Error().Err(err).Msg("cant request getting chats")
		return nil
	}

	return resp
}
