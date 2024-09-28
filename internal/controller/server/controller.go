package handlers

import (
	"messenger/internal/interfaces"
	"messenger/internal/service/service"
)

type Controller struct {
	repo interfaces.Repository
	cs   *service.ChatService
}

func New(repo interfaces.Repository, cs *service.ChatService) *Controller {
	return &Controller{repo: repo, cs: cs}
}
