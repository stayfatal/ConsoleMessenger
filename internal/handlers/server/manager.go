package handlers

import (
	"messenger/internal/database"
)

type handlersManager struct {
	dm *database.DbManager
}

func NewHandlersManager() *handlersManager {
	dm := database.GetDbManager()
	return &handlersManager{dm: dm}
}
