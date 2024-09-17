package database

type Chat struct {
	Id   int
	Name string
}

func (dm *DbManager) CreateChat(chat Chat) (id int, err error) {
	err = dm.db.QueryRow("insert into chats (name) values ($1) returning id", chat.Name).Scan(&id)
	return
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
