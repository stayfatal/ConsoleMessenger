package websocket

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/gobwas/ws/wsutil"
	"github.com/rs/zerolog/log"
)

func Reader(wg *sync.WaitGroup, conn net.Conn) {
	defer wg.Done()
	for {
		msg, err := wsutil.ReadServerText(conn)
		if err != nil {
			log.Error().Err(err).Msg("cant read message from conn")
			return
		}

		fmt.Println(string(msg))
	}
}

func Writer(wg *sync.WaitGroup, conn net.Conn) {
	defer wg.Done()
	r := bufio.NewReader(os.Stdin)
	for {
		r.Reset(os.Stdin)

		msg, err := r.ReadString('\n')
		if err != nil {
			log.Error().Err(err).Msg("cant read message from console")
			return
		}
		fmt.Println()

		err = wsutil.WriteClientText(conn, []byte(msg))
		if err != nil {
			log.Error().Err(err).Msg("cant write message to conn")
			return
		}
	}
}
