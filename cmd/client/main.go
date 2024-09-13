package main

import (
	"messenger/internal/iface"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	im := iface.NewInterfaceManager()

	log.Info().Msg("client is running")
	im.RunApp()
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
