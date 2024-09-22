package models

import "time"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Message struct {
	Id      int
	ChatId  int
	UserId  int
	Date    time.Time
	Message string
}

type ShortMessage struct {
	SenderUsername string `json:"sender_username"`
	Message        string `json:"message"`
}

type Chat struct {
	Id   int
	Name string
}

type ChatMember struct {
	Id     int
	ChatId int
	UserId int
}
