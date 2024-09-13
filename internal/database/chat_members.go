package database

type ChatMember struct {
	Id     int
	ChatId int
	UserId int
}

func (dm *DbManager) AddChatMember(cm ChatMember) error {
	_, err := dm.db.Exec("insert into chat_members (chat_id,user_id) values ($1,$2)", cm.ChatId, cm.UserId)
	return err
}

func (dm *DbManager) GetAllUserChats(userId int) (map[int]string, error) {
	req := `SELECT chat_members.chat_id, chats.name 
	FROM chats 
	JOIN chat_members ON chat_members.chat_id = chats.id 
	WHERE chat_members.user_id = $1`
	rows, err := dm.db.Query(req, userId)
	if err != nil {
		return nil, err
	}

	result := make(map[int]string)
	for rows.Next() {
		var (
			id   int
			name string
		)
		err = rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}

		result[id] = name
	}

	return result, nil
}

func (dm *DbManager) GetAllChatMembers(chatId int) ([]int, error) {
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
