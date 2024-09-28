package controller

import (
	"context"
	"fmt"
	chatSvc "messenger/internal/service/client"
	"messenger/internal/utils"
	"sync"

	"github.com/gobwas/ws"
	"github.com/rs/zerolog/log"
)

func (hm *HandlersManager) NewChatHandler(recipient string) {
	conn, _, _, err := ws.Dialer{
		Header: ws.HandshakeHeaderHTTP{
			"Authorization": []string{utils.GetToken()},
			"Recipient":     []string{recipient},
		},
	}.Dial(context.TODO(), fmt.Sprintf("wss://%s/ws/chats", hm.addr))

	if err != nil {
		log.Error().Err(err).Msg("cant request create new chat")
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go chatSvc.Reader(&wg, conn)
	go chatSvc.Writer(&wg, conn)
	wg.Wait()
}

func (hm *HandlersManager) JoinChatHandler(chatId string) {
	conn, _, _, err := ws.Dialer{
		Header: ws.HandshakeHeaderHTTP{
			"Authorization": []string{utils.GetToken()},
		},
	}.Dial(context.TODO(), fmt.Sprintf("wss://%s/ws/chats/%s", hm.addr, chatId))

	if err != nil {
		log.Error().Err(err).Msg("cant request join chat")
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go chatSvc.Reader(&wg, conn)
	go chatSvc.Writer(&wg, conn)
	wg.Wait()
}
