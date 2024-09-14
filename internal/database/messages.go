package database

import "time"

type Message struct {
	Id     int
	ChatId int
	UserId int
	Date   time.Time
	Text   string
}

func (dm *DbManager) SaveMessage(msg Message) error {
	_, err := dm.db.Exec("insert into messages (chat_id,user_id,message) values ($1,$2,$3)", msg.ChatId, msg.UserId, msg.Text)
	return err
}
