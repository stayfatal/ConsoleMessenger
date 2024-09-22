package iface

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

func (im *InterfaceManager) switchToStartMenu() {
	str := `Register Menu
1) Register
2) Login
3) Quit`
	fmt.Println(str)

	var option int
	for {
		_, err := fmt.Scanf("%d\n", &option)
		if err != nil {
			log.Error().Err(err).Msg("incorrect input")
			continue
		}

		switch option {
		case 1:
			im.switchToRegistrationMenu()
		case 2:
			im.switchToLoginMenu()
		case 3:
			os.Exit(0)
		default:
			log.Info().Msg(fmt.Sprintf("there is no option %d)", option))
			continue
		}
		break
	}
}

func (im *InterfaceManager) switchToRegistrationMenu() {
	fmt.Println("Please enter login and password in format login password")

	var username, password string
	for {
		_, err := fmt.Scanf("%s %s\n", &username, &password)
		if err != nil {
			log.Error().Err(err).Msg("incorrect input please check correct format")
			continue
		}
		break
	}

	im.hm.RegistrationHandler(username, password)
	im.switchToMainMenu()
}

func (im *InterfaceManager) switchToLoginMenu() {
	fmt.Println("Please enter login and password in format login password")

	var username, password string
	for {
		_, err := fmt.Scanf("%s %s\n", &username, &password)
		if err != nil {
			log.Error().Err(err).Msg("incorrect input please check correct format")
			continue
		}
		break
	}

	im.hm.LoginHandler(username, password)
	im.switchToMainMenu()
}

func (im *InterfaceManager) switchToChatsMenu() {
	str := "Chats Menu\n1) New chat\n"
	chats := im.hm.GetChatsHandler()
	options := make(map[int]string)
	var i int = 2
	for id, name := range chats {
		str += fmt.Sprintf("%d) Chat : %s\n", i, name)
		options[i] = id
		i++
	}
	str += fmt.Sprintf("%d) Quit", i)
	fmt.Println(str)

	var option int
	for {
		_, err := fmt.Scanf("%d\n", &option)
		if err != nil {
			log.Error().Err(err).Msg("incorrect input")
			continue
		}

		switch option {
		case 1:
			im.switchToChatCreationChatMenu()
		case i:
			os.Exit(0)
		default:
			if option > i-1 || option < 1 {
				log.Info().Msg(fmt.Sprintf("there is no option %d)", option))
				continue
			}
			im.switchToJoinChatMenu(options[option])
		}
		break
	}
}

func (im *InterfaceManager) switchToChatCreationChatMenu() {
	fmt.Println("Please enter recipient's username\nTo exit write \"/back\"")

	var username string
	for {
		_, err := fmt.Scanf("%s\n", &username)
		if err != nil {
			log.Error().Err(err).Msg("incorrect input")
			continue
		}

		if username == "/back" {
			im.switchToMainMenu()
		}
		break
	}

	im.switchToNewChatMenu(username)
}

func (im *InterfaceManager) switchToNewChatMenu(username string) {
	str := fmt.Sprintf("Chat with %s", username)
	fmt.Println(str)

	im.hm.NewChatHandler(username)
}

func (im *InterfaceManager) switchToJoinChatMenu(chatId string) {
	messages := im.hm.ChatHistoryHandler(chatId)
	for i := len(messages) - 1; i >= 0; i-- {
		fmt.Printf("%s : %s\n", messages[i].SenderUsername, messages[i].Message)
	}
	im.hm.JoinChatHandler(chatId)
}

func (im *InterfaceManager) switchToMainMenu() {
	if im.hm.ValidateTokenHandler() {
		im.switchToChatsMenu()
	} else {
		im.switchToStartMenu()
	}
}
