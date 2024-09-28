package repository

import (
	"messenger/internal/models"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func (dm *repository) CreateUser(user models.User) (int, error) {
	var id int
	err := dm.db.QueryRow("insert into users (username,password) values ($1,$2) returning id", user.Username, user.Password).Scan(&id)
	return id, errors.Wrap(err, "inserting into users")
}

func (dm *repository) GetUserByName(username string) (models.User, error) {
	scannedUser := models.User{}
	err := dm.db.QueryRow("select * from users where username = $1", username).Scan(&scannedUser.Id, &scannedUser.Username, &scannedUser.Password)
	return scannedUser, errors.Wrap(err, "selecting from users by username")
}

func (dm *repository) GetUserIdByName(username string) (id int, err error) {
	err = dm.db.QueryRow("select id from users where username = $1", username).Scan(&id)
	return id, errors.Wrap(err, "selecting id from users users by username")
}

func (dm *repository) GetUserById(id int) (models.User, error) {
	user := models.User{}
	err := dm.db.QueryRow("select * from users where id = $1", id).Scan(&user.Id, &user.Username, &user.Password)
	return user, errors.Wrap(err, "selecting from users by id")
}

func (dm *repository) GetUsernameById(id int) (string, error) {
	var username string
	err := dm.db.QueryRow("select username from users where id = $1", id).Scan(&username)
	return username, errors.Wrap(err, "selecting username from users by id")
}
