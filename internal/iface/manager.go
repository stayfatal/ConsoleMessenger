package iface

import (
	handlers "messenger/internal/handlers/client"

	"github.com/rivo/tview"
)

type InterfaceManager struct {
	app *tview.Application
	hm  *handlers.HandlersManager
}

func NewInterfaceManager() *InterfaceManager {
	return &InterfaceManager{app: tview.NewApplication(), hm: handlers.GetHandlersManager()}
}

func (im *InterfaceManager) RunApp() error {
	im.switchToMainMenu()
	return im.app.Run()
}
