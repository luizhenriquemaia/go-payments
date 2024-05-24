package configs

import (
	"database/sql"
	"errors"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("record doesn't exists")
	ErrCreateFailed = errors.New("create failed")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type Repository struct {
	db *sql.DB
}

func New_repository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Migrate() error {
	query := ``
	_, err := r.db.Exec(query)
	return err
}
