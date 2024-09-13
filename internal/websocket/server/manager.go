package websocket

import (
	"messenger/internal/database"
	"net"
	"sync"
)

var (
	once sync.Once
	wm   *WebsocketManager
)

type WebsocketManager struct {
	dm          *database.DbManager
	chatMembers map[int]*ChatMember
	mu          sync.Mutex
}

type ChatMember struct {
	conn net.Conn
	out  chan []byte
}

func GetWebsocketManager() *WebsocketManager {
	once.Do(func() {
		wm = &WebsocketManager{chatMembers: make(map[int]*ChatMember), dm: database.GetDbManager()}
	})

	return wm
}
