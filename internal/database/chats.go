package database

type Chat struct {
	Id   int
	Name string
}

func (dm *DbManager) CreateChat(chat Chat) (id int, err error) {
	err = dm.db.QueryRow("insert into chats (name) values ($1) returning id", chat.Name).Scan(&id)
	return
}
