package service

import "net"

func (cs *ChatService) addChatMember(id int, conn net.Conn) {
	cs.mu.Lock()
	cs.chatMembers[id] = &chatMember{conn: conn, out: make(chan []byte, 5)}
	cs.mu.Unlock()
}

func (cs *ChatService) getChatMember(id int) *chatMember {
	cs.mu.Lock()
	cm := cs.chatMembers[id]
	cs.mu.Unlock()
	return cm
}

func (cs *ChatService) deleteChatMember(id int) {
	cs.mu.Lock()
	if _, ok := cs.chatMembers[id]; ok {
		cs.chatMembers[id].conn.Close()
		close(cs.chatMembers[id].out)
		delete(cs.chatMembers, id)
	}
	cs.mu.Unlock()
}

func (cs *ChatService) isConnected(id int) bool {
	cs.mu.Lock()
	_, ok := cs.chatMembers[id]
	cs.mu.Unlock()
	return ok
}
