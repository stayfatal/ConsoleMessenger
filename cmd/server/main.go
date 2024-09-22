package main

import (
	handlers "messenger/internal/handlers/server"
	"messenger/internal/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	router := gin.Default()

	hm := handlers.NewHandlersManager()

	router.POST("/register", hm.RegisterHandler)
	router.POST("/login", hm.LoginHandler)

	auth := router.Group("/")
	auth.Use(middleware.Authentication())

	auth.GET("/chats", hm.GetChatsHandler)
	auth.GET("/chats/:id/history", hm.GetLastChatMessagesHandler)
	auth.GET("/token", hm.ValidateTokenHandler)

	auth.GET("/ws/chats", hm.NewChatHandler)
	auth.GET("/ws/chats/:id", hm.JoinChatHandler)

	if err := router.Run(":80"); err != nil {
		log.Fatal().Err(err).Msg("Failed to run server")
	}
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}
