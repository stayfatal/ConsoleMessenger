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
}

func GetHandlersManager() *HandlersManager {
	once.Do(func() {
		client := resty.New()
		hm = &HandlersManager{client: client}
	})

	return hm
}
