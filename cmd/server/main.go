package main

import (
	"fmt"
	"messenger/config"
	controller "messenger/internal/controller/server"
	"messenger/internal/middleware"
	"messenger/internal/repository"
	"messenger/internal/service/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	router := gin.Default()

	cfg := config.New()

	db := config.NewDB(cfg)

	repo := repository.New(db)

	cs := service.New(repo)

	cr := controller.New(repo, cs)

	router.POST("/register", cr.Register)
	router.POST("/login", cr.Login)

	auth := router.Group("/")
	auth.Use(middleware.Authentication())

	auth.GET("/chats", cr.GetChats)
	auth.GET("/chats/:id/history", cr.GetLastChatMessages)
	auth.GET("/token", cr.ValidateTokenHandler)

	auth.GET("/ws/chats", cr.NewChat)
	auth.GET("/ws/chats/:id", cr.JoinChat)

	if err := router.Run(fmt.Sprintf(":%s", cfg.PORT)); err != nil {
		log.Fatal().Err(err).Msg("Failed to run server")
	}
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}
