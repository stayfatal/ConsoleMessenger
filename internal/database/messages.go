package database

import "github.com/google/uuid"

type Message struct {
	ChatId   uuid.UUID
	Messages []string
}

func SaveMessages() {

}
