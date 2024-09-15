package handlers

import (
	"sync"

	"github.com/go-resty/resty/v2"
)

var (
	hm   *HandlersManager
	once sync.Once
)

type HandlersManager struct {
	client *resty.Client
	addr   string
}

func GetHandlersManager() *HandlersManager {
	once.Do(func() {
		client := resty.New()
		hm = &HandlersManager{client: client, addr: "df7a-2a00-1370-81aa-27da-1932-6373-3bb2-170e.ngrok-free.app"}
	})

	return hm
}
