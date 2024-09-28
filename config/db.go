package config

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func NewDB(cfg *Config) *sqlx.DB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_SSL_MODE)
	db, err := sqlx.Open(cfg.DB_DRIVER, connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("cant connect db")
	}

	if err := db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("cant ping db")

	}

	return db
}
