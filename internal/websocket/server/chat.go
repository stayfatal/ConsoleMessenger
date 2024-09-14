package websocket

import (
	"fmt"
	"messenger/internal/database"

	"github.com/gobwas/ws/wsutil"
	"github.com/rs/zerolog/log"
)

func (wm *WebsocketManager) StartChat(chatId int) {
	members, err := wm.dm.GetAllChatMembers(chatId)
	if err != nil {
		log.Error().Err(err).Msg("cant get all chat's members")
		return
	}

	for _, id := range members {
		if wm.IsConnected(id) {
			wm.StartReader(id, chatId, members)
			wm.StartWriter(id)
		}
	}
}

func (wm *WebsocketManager) StartReader(id, chatId int, members []int) {
	go func() {
		cm := wm.GetChatMember(id)
		for {
			msg, _, err := wsutil.ReadClientData(cm.conn)
			if err != nil {
				log.Error().Err(err).Msg("cant read from conn")
				return
			}
			fmt.Println(string(msg))

			wm.BroadCast(id, chatId, msg, members)
		}
	}()
}

func (wm *WebsocketManager) StartWriter(id int) {
	go func() {
		cm := wm.GetChatMember(id)
		for {
			msg := <-cm.out
			err := wsutil.WriteServerText(cm.conn, msg)
			if err != nil {
				log.Error().Err(err).Msg("cant write to conn")
				return
			}
		}
	}()
}

func (wm *WebsocketManager) BroadCast(id, chatId int, msg []byte, members []int) {
	for _, recipientId := range members {
		if recipientId != id && wm.IsConnected(recipientId) {
			recipient := wm.GetChatMember(recipientId)
			recipient.out <- msg
		}

		go wm.dm.SaveMessage(database.Message{
			ChatId: chatId,
			UserId: id,
			Text:   string(msg),
		})
	}
}
