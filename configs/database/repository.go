package database

import (
	"database/sql"
	"embed"
	"errors"

	"github.com/pressly/goose/v3"
)

var (
	ErrMigration    = errors.New("migration error - details: ")
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("record doesn't exists")
	ErrCreateFailed = errors.New("create failed")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Migrate(file_sys embed.FS) error {
	goose.SetBaseFS(file_sys)

	err := goose.SetDialect("postgres")
	if err != nil {
		return errors.New(ErrMigration.Error() + " fase 1 " + err.Error())
	}

	err = goose.Up(r.db, "configs/database/migrations")
	if err != nil {
		return errors.New(ErrMigration.Error() + " fase 2 - " + err.Error())
	}

	return err
}
