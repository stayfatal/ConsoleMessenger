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
	_, err := dm.db.Exec("insert into messages (sender_id,recipient_id,message) values ($1,$2,$3)", msg.ChatId, msg.UserId, msg.Text)
	return err
}
