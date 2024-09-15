package handlers

import (
	"fmt"
	"messenger/internal/env"

	"github.com/rs/zerolog/log"
)

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
