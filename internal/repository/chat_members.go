package repository

import "messenger/internal/models"

func (dm *repository) AddChatMember(cm models.ChatMember) error {
	_, err := dm.db.Exec("insert into chat_members (chat_id,user_id) values ($1,$2)", cm.ChatId, cm.UserId)
	return err
}

func (dm *repository) GetAllChatMembers(chatId int) ([]int, error) {
	rows, err := dm.db.Query("select user_id from chat_members where chat_id = $1", chatId)
	if err != nil {
		return nil, err
	}

	var result []int
	for rows.Next() {
		var userId int
		err := rows.Scan(&userId)
		if err != nil {
			return nil, err
		}

		result = append(result, userId)
	}

	return result, nil
}
