package main

import (
	handlers "messenger/internal/handlers/server"
	"messenger/internal/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	router := gin.Default()

	hm := handlers.NewHandlersManager()

	router.POST("/register", hm.CreateUserHandler)
	router.POST("/login", hm.LoginHandler)

	auth := router.Group("/")
	auth.Use(middleware.Authentication())

	auth.GET("/chats", hm.GetChatsHandler)
	auth.GET("/token", hm.ValidateTokenHandler)

	auth.GET("/ws/chats", hm.NewChatHandler)
	auth.GET("/ws/chats/:chat_id", hm.JoinChatHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal().Err(err).Msg("Failed to run server")
	}
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
