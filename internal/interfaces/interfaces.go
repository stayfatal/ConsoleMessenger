package interfaces

import (
	"messenger/internal/models"
)

type Repository interface {
	CreateUser(user models.User) (int, error)
	GetUserByName(username string) (models.User, error)
	GetUserIdByName(username string) (id int, err error)
	GetUserById(id int) (models.User, error)
	GetUsernameById(id int) (string, error)

	CreateChat(chat models.Chat) (id int, err error)
	GetAllUserChats(userId int) (map[int]string, error)

	AddChatMember(cm models.ChatMember) error
	GetAllChatMembers(chatId int) ([]int, error)
	SaveMessage(msg models.Message) error
	GetLastChatMessages(chatId int) ([]models.ShortMessage, error)
}
