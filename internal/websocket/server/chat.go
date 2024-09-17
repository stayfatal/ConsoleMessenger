package websocket

import (
	"messenger/internal/database"

	"github.com/gobwas/ws/wsutil"
	"github.com/rs/zerolog/log"
)

func (wm *WebsocketManager) JoinChat(id, chatId int) {
	wm.StartReader(id, chatId)
	wm.StartWriter(id)
}

func (wm *WebsocketManager) StartReader(id, chatId int) {
	go func() {
		cm := wm.GetChatMember(id)
		for {
			msg, err := wsutil.ReadClientText(cm.conn)
			if err != nil {
				wm.DeleteChatMember(id)
				log.Error().Err(err).Msg("cant read from conn")
				return
			}

			wm.BroadCast(id, chatId, msg)

			errCh := make(chan error)
			go func(errCh chan<- error) {
				err := wm.dm.SaveMessage(database.Message{
					ChatId:  chatId,
					UserId:  id,
					Message: string(msg),
				})

				errCh <- err
			}(errCh)

			err = <-errCh
			if err != nil {
				log.Error().Err(err).Msg("cant save message")
			}
		}
	}()
}

func (wm *WebsocketManager) StartWriter(id int) {
	go func() {
		cm := wm.GetChatMember(id)
		// type Message struct {
		// 	senderUsername string
		// 	message        string
		// }
		// senderUsername, err := wm.dm.GetUsernameById(id)
		// if err != nil {
		// 	log.Error().Err(err).Msg("cant get username from db")
		// 	return
		// }
		for {
			// message := Message{senderUsername: senderUsername}

			msg := <-cm.out
			// message.message=string(msg)

			err := wsutil.WriteServerText(cm.conn, msg)
			if err != nil {
				wm.DeleteChatMember(id)
				log.Error().Err(err).Msg("cant write to conn")
				return
			}
		}
	}()
}

func (wm *WebsocketManager) BroadCast(id, chatId int, msg []byte) {
	members, err := wm.dm.GetAllChatMembers(chatId)
	if err != nil {
		log.Error().Err(err).Msg("cant get all chat's members")
		return
	}
	for _, recipientId := range members {
		if recipientId != id && wm.IsConnected(recipientId) {
			recipient := wm.GetChatMember(recipientId)
			recipient.out <- msg
		}
	}
}
