package websocket

import (
	"net"
	"sync"

	"github.com/google/uuid"
)

var (
	mu    sync.Mutex
	conns map[uuid.UUID]net.Conn = make(map[uuid.UUID]net.Conn)
)

func AddConn(id uuid.UUID, conn net.Conn) {
	mu.Lock()
	conns[id] = conn
	mu.Unlock()
}

func GetConn(id uuid.UUID) net.Conn {
	mu.Lock()
	conn := conns[id]
	mu.Unlock()
	return conn
}
