package iface

import (
	handlers "messenger/internal/handlers/client"
)

type InterfaceManager struct {
	hm *handlers.HandlersManager
}

func NewInterfaceManager() *InterfaceManager {
	return &InterfaceManager{hm: handlers.GetHandlersManager()}
}

func (im *InterfaceManager) RunApp() {
	im.switchToMainMenu()
}
