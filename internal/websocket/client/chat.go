package websocket

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/rs/zerolog/log"
)

func Reader(conn net.Conn) {
	for {
		msg, _, err := wsutil.ReadServerData(conn)

		if err != nil {
			log.Error().Err(err).Msg("cant read message from server")
			return
		}

		fmt.Println(string(msg))
	}
}

func Writer(conn net.Conn) {
	for {
		in := bufio.NewReader(os.Stdin)

		var msg []byte
		fmt.Fscanln(in, &msg)

		err := wsutil.WriteServerMessage(conn, ws.OpBinary, msg)
		if err != nil {
			log.Error().Err(err).Msg("cant write message to server")
			return
		}
	}
}
