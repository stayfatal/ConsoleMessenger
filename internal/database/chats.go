package database

import (
	"fmt"

	"github.com/google/uuid"
)

type Chat struct {
	Id      uuid.UUID
	User1Id uuid.UUID
	User2Id uuid.UUID
}

func (dm *DbManager) CreateChat(chat Chat) error {
	_, err := dm.db.Exec("insert into chats (id,user1_id,user2_id) values ($1,$2,$3)", chat.Id, chat.User1Id, chat.User2Id)
	return err
}

func (dm *DbManager) GetChatById(id uuid.UUID) (Chat, error) {
	chat := Chat{}
	err := dm.db.QueryRow("select * from chats where id = $1", id).Scan(&chat.Id, &chat.User1Id, &chat.User2Id)
	return chat, err
}

func (dm *DbManager) GetAllUserChats(id uuid.UUID) (map[string]string, error) {
	chats := make(map[string]string)
	rows, err := dm.db.Query("select * from chats where user1_id = $1 or user2_id = $1", id)
	if err != nil {
		return nil, err
	}

	var i int
	for rows.Next() {
		var chat Chat
		rows.Scan(&chat.Id, &chat.User1Id, &chat.User2Id)
		chats[fmt.Sprintf("Chat %d", i)] = chat.Id.String()
		i++
	}

	return chats, nil
}
