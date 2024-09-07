package handlers

import (
	"context"
	"fmt"
	"messenger/internal/env"
	websocket "messenger/internal/websocket/client"

	"github.com/gobwas/ws"
	"github.com/rs/zerolog/log"
)

func (hm *HandlersManager) NewChatHandler(recipient string) {
	conn, _, _, err := ws.Dialer{
		Header: ws.HandshakeHeaderHTTP{
			"Authorization": []string{env.GetToken()},
			"Recipient":     []string{recipient},
		},
	}.Dial(context.TODO(), "ws://localhost:8080/ws/chats")

	if err != nil {
		log.Error().Err(err).Msg("cant request create new chat")
		return
	}

	go websocket.Reader(conn)
	go websocket.Writer(conn)
}

func (hm *HandlersManager) JoinChatHandler(chatId string) {
	conn, _, _, err := ws.Dialer{
		Header: ws.HandshakeHeaderHTTP{
			"Authorization": []string{env.GetToken()},
		},
	}.Dial(context.TODO(), fmt.Sprintf("ws://localhost:8080/ws/chats/:%s", chatId))

	if err != nil {
		log.Error().Err(err).Msg("cant request join chat")
		return
	}

	go websocket.Reader(conn)
	go websocket.Writer(conn)
}
