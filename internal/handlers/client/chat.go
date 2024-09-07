package handlers

import (
	"messenger/internal/env"

	"github.com/rs/zerolog/log"
)

func (hm *HandlersManager) GetChatsHandler() map[string]string {
	resp := make(map[string]string)

	_, err := hm.client.R().
		SetHeader("Authorization", env.GetToken()).
		SetResult(&resp).
		Get("http://localhost:8080/chats")

	if err != nil {
		log.Error().Err(err).Msg("cant request getting chats")
		return nil
	}

	return resp
}
