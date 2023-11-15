package storage

import "github.com/jmoiron/sqlx"

type Repository struct {
}

func New(db *sqlx.DB) *Repository {
	return &Repository{}
}
