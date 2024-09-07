package websocket

import (
	"messenger/internal/database"
	"net"

	"github.com/gobwas/ws/wsutil"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type ChatRoom struct {
	database.Chat
	conn1 net.Conn
	conn2 net.Conn
}

func StartChat(id1, id2 uuid.UUID) {
	conn1 := GetConn(id1)
	conn2 := GetConn(id2)

	handleConn(conn1, conn2)
	handleConn(conn2, conn1)
}

func handleConn(sender, recipient net.Conn) {
	for {
		msg, op, err := wsutil.ReadClientData(sender)
		if err != nil {
			log.Error().Err(err).Msg("cant read message from sender")
			return
		}

		err = wsutil.WriteClientMessage(recipient, op, msg)
		if err != nil {
			log.Error().Err(err).Msg("cant write message to recipient")
			return
		}
	}
}
