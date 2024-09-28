package service

import (
	"messenger/internal/interfaces"
	"net"
	"sync"
)

type ChatService struct {
	repo        interfaces.Repository
	chatMembers map[int]*chatMember
	mu          sync.Mutex
}

type chatMember struct {
	conn net.Conn
	out  chan []byte
}

func New(repo interfaces.Repository) *ChatService {
	return &ChatService{repo: repo, chatMembers: make(map[int]*chatMember)}
}
