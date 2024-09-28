package service

import (
	"messenger/internal/models"

	"github.com/gobwas/ws/wsutil"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (cs *ChatService) JoinChat(id, chatId int) {
	cs.startReader(id, chatId)
	cs.startWriter(id)
}

func (wm *ChatService) startReader(id, chatId int) {
	go func() {
		cm := wm.getChatMember(id)
		for {
			msg, err := wsutil.ReadClientText(cm.conn)
			if err != nil {
				wm.deleteChatMember(id)
				log.Error().Stack().Err(errors.Wrap(err, "reading client message")).Msg("")
				return
			}

			wm.broadcast(id, chatId, msg)

			errCh := make(chan error)
			go func(errCh chan<- error) {
				err := wm.repo.SaveMessage(models.Message{
					ChatId:  chatId,
					UserId:  id,
					Message: string(msg),
				})

				errCh <- err
			}(errCh)

			err = <-errCh
			if err != nil {
				log.Error().Stack().Err(errors.Wrap(err, "saving message to db")).Msg("")
				return
			}
		}
	}()
}

func (wm *ChatService) startWriter(id int) {
	go func() {
		cm := wm.getChatMember(id)
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
				wm.deleteChatMember(id)
				log.Error().Stack().Err(errors.Wrap(err, "writing message to client")).Msg("")
				return
			}
		}
	}()
}

func (wm *ChatService) broadcast(id, chatId int, msg []byte) {
	members, err := wm.repo.GetAllChatMembers(chatId)
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling GetAllChatMembers")).Msg("")
		return
	}
	for _, recipientId := range members {
		if recipientId != id && wm.isConnected(recipientId) {
			recipient := wm.getChatMember(recipientId)
			recipient.out <- msg
		}
	}
}
