package iface

import (
	controller "messenger/internal/controller/client"
)

type InterfaceManager struct {
	hm *controller.HandlersManager
}

func NewInterfaceManager() *InterfaceManager {
	return &InterfaceManager{hm: controller.GetHandlersManager()}
}

func (im *InterfaceManager) RunApp() {
	im.switchToMainMenu()
}
