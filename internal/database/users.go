package database

import (
	_ "github.com/lib/pq"
)

type User struct {
	Id       int
	Username string `json:"username"`
	Password string `json:"password"`
}

func (dm *DbManager) CreateUser(user User) (id int, err error) {
	err = dm.db.QueryRow("insert into users (username,password) values ($1,$2) returning id", user.Username, user.Password).Scan(&id)
	return
}

func (dm *DbManager) GetUserByName(username string) (User, error) {
	scannedUser := User{}
	err := dm.db.QueryRow("select * from users where username = $1", username).Scan(&scannedUser.Id, &scannedUser.Username, &scannedUser.Password)
	return scannedUser, err
}

func (dm *DbManager) GetUserIdByName(username string) (id int, err error) {
	err = dm.db.QueryRow("select id from users where username = $1", username).Scan(&id)
	return id, err
}

func (dm *DbManager) GetUserById(id int) (user User, err error) {
	err = dm.db.QueryRow("select * from users where id = $1", id).Scan(&user.Id, &user.Username, &user.Password)
	return
}
