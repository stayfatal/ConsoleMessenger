package main

import (
	"messenger/internal/iface"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	im := iface.
		NewInterfaceManager()

	if err := im.RunApp(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
