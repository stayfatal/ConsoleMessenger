package database

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type DbManager struct {
	db *sql.DB
}

var (
	dm   *DbManager
	once sync.Once
)

func GetDbManager() *DbManager {
	once.Do(func() {
		connStr := "user=postgres password=mypass dbname=messenger sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal().Err(err).Msg("cant connect to database")
		}
		dm = &DbManager{db: db}
	})

	return dm
}
