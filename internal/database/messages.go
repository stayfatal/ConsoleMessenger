package database

import (
	"messenger/internal/models"
)

func (dm *DbManager) SaveMessage(msg models.Message) error {
	_, err := dm.db.Exec("insert into messages (chat_id,user_id,message) values ($1,$2,$3)", msg.ChatId, msg.UserId, msg.Message)
	return err
}

func (dm *DbManager) GetLastChatMessages(chatId int) ([]models.ShortMessage, error) {
	req := `SELECT users.username, messages.message
	FROM messages
	JOIN users ON messages.user_id = users.id
	WHERE messages.chat_id = $1
	ORDER BY messages.id DESC
	LIMIT 20`
	rows, err := dm.db.Query(req, chatId)
	if err != nil {
		return nil, err
	}

	var messages []models.ShortMessage
	for rows.Next() {
		var senderUsername, msg string
		err = rows.Scan(&senderUsername, &msg)
		if err != nil {
			return nil, err
		}

		messages = append(messages, models.ShortMessage{SenderUsername: senderUsername, Message: msg})
	}

	return messages, nil
}
