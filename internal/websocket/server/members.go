package websocket

import "net"

func (wm *WebsocketManager) AddChatMember(id int, conn net.Conn) {
	wm.mu.Lock()
	wm.chatMembers[id] = &ChatMember{conn: conn, out: make(chan []byte, 5)}
	wm.mu.Unlock()
}

func (wm *WebsocketManager) GetChatMember(id int) *ChatMember {
	wm.mu.Lock()
	cm := wm.chatMembers[id]
	wm.mu.Unlock()
	return cm
}

func (wm *WebsocketManager) DeleteChatMember(id int) {
	wm.mu.Lock()
	wm.chatMembers[id].conn.Close()
	close(wm.chatMembers[id].out)
	delete(wm.chatMembers, id)
	wm.mu.Unlock()
}

func (wm *WebsocketManager) IsConnected(id int) bool {
	wm.mu.Lock()
	_, ok := wm.chatMembers[id]
	wm.mu.Unlock()
	return ok
}
