package database

import (
	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

type User struct {
	Id       uuid.UUID
	Username string `json:"username"`
	Password string `json:"password"`
}

func (dm *DbManager) CreateUser(user User) error {
	_, err := dm.db.Exec("insert into users (id,username,password) values ($1,$2,$3)", user.Id, user.Username, user.Password)
	return err
}

func (dm *DbManager) GetUserByName(username string) (User, error) {
	scannedUser := User{}
	err := dm.db.QueryRow("select * from users where username = $1", username).Scan(&scannedUser.Id, &scannedUser.Username, &scannedUser.Password)
	return scannedUser, err
}
