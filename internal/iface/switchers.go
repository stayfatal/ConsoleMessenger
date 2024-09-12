package iface

import (
	"os"

	"github.com/rivo/tview"
)

func (im *InterfaceManager) switchToStartMenu() {
	startMenu := tview.NewList().
		AddItem("Register", "", '1', func() {
			im.switchToRegistrationMenu()
		}).
		AddItem("Login", "", '2', func() {
			im.switchToLoginMenu()
		}).
		AddItem("Quit", "Exit from app", '3', func() {
			im.app.Stop()
			os.Exit(0)
		})
	im.app.SetRoot(startMenu, true).SetFocus(startMenu)
}

func (im *InterfaceManager) switchToChatsMenu() {
	chatsMenu := tview.NewList().
		AddItem("New chat", " ", '1', func() {
			im.switchToNewChatMenu()
		})

	chats := im.hm.GetChatsHandler()

	for key, val := range chats {
		chatsMenu.AddItem(val, key, ' ', func() {
			im.app.Stop()
		})
	}

	chatsMenu.AddItem("Quit", "Exit from app", '2', func() {
		im.app.Stop()
		os.Exit(0)
	})
	im.app.SetRoot(chatsMenu, true).SetFocus(chatsMenu)
}

func (im *InterfaceManager) switchToNewChatMenu() {
	var username string
	newChatMenu := tview.NewForm().
		AddInputField("username : ", "", 20, nil, func(text string) {
			username = text
		}).
		AddButton("Enter", func() {
			im.hm.NewChatHandler(username)
			im.switchToMainMenu()
		}).
		AddButton("Back", im.switchToStartMenu)
	im.app.SetRoot(newChatMenu, true).SetFocus(newChatMenu)
}

func (im *InterfaceManager) switchToRegistrationMenu() {
	var username, password string
	registrationMenu := tview.NewForm().
		AddInputField("Username: ", "", 20, nil, func(text string) {
			username = text
		}).
		AddInputField("Password : ", "", 20, nil, func(text string) {
			password = text
		}).
		AddButton("Enter", func() {
			im.hm.RegistrationHandler(username, password)
			im.switchToMainMenu()
		}).
		AddButton("Back", im.switchToStartMenu)
	im.app.SetRoot(registrationMenu, true).SetFocus(registrationMenu)
}

func (im *InterfaceManager) switchToLoginMenu() {
	var username, password string
	loginMenu := tview.NewForm().
		AddInputField("Username: ", "", 20, nil, func(text string) {
			username = text
		}).
		AddInputField("Password : ", "", 20, nil, func(text string) {
			password = text
		}).
		AddButton("Enter", func() {
			im.hm.LoginHandler(username, password)
			im.switchToMainMenu()
		}).AddButton("Back", im.switchToStartMenu)
	im.app.SetRoot(loginMenu, true).SetFocus(loginMenu)
}

func (im *InterfaceManager) switchToMainMenu() {
	if im.hm.ValidateTokenHandler() {
		im.switchToChatsMenu()
	} else {
		im.switchToStartMenu()
	}
}
