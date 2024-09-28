package main

import (
	"fmt"
	"messenger/config"
	controller "messenger/internal/controller/server"
	"messenger/internal/middleware"
	"messenger/internal/repository"
	"messenger/internal/service/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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

	errs := make(chan error)
	go func() {
		err := router.Run(fmt.Sprintf(":%s", cfg.PORT))
		errs <- errors.Wrap(err, "running server")
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		buf := <-c
		errs <- errors.New(buf.String())
	}()

	log.Fatal().Err(<-errs).Msg("")
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}
