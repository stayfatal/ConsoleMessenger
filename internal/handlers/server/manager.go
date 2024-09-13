package handlers

import (
	"messenger/internal/database"
	websocket "messenger/internal/websocket/server"
)

type handlersManager struct {
	dm *database.DbManager
	wm *websocket.WebsocketManager
}

func NewHandlersManager() *handlersManager {
	return &handlersManager{dm: database.GetDbManager(), wm: websocket.GetWebsocketManager()}
}
