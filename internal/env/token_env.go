package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func WriteToken(token string) {
	envFile := "token.env"

	file, err := os.OpenFile(envFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Error().Err(err).Msg("")
		return
	}
	defer file.Close()

	envLine := fmt.Sprintf("REGISTRATION_TOKEN=%s\n", token)

	if _, err := file.WriteString(envLine); err != nil {
		log.Error().Err(err).Msg("cant write to token.env")
		return
	}
}

func GetToken() string {
	err := godotenv.Load("token.env")
	if err != nil {
		log.Error().Err(err).Msg("cant load token.env")
		return ""
	}
	return os.Getenv("REGISTRATION_TOKEN")
}
