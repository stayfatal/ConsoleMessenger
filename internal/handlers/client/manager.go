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
		hm = &HandlersManager{client: client, addr: "genuine-fish-light.ngrok-free.app"}
	})

	return hm
}
