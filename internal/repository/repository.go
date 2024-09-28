package repository

import (
	"messenger/internal/interfaces"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) interfaces.Repository {
	return &repository{db: db}
}
