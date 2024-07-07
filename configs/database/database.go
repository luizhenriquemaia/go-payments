package database

import (
	"database/sql"
	"embed"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func GetDb() *sql.DB {
	connection_str := os.Getenv("POSTGRES_CONNECTION")

	if connection_str == "" {
		connection_str = "postgresql://default:password123@localhost:5433/go-payments?sslmode=disable"
	}

	db, err := sql.Open("postgres", connection_str)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InitDb(file_sys embed.FS) {
	db := GetDb()
	repository := NewRepository(db)

	if err := repository.Migrate(file_sys); err != nil {
		log.Fatal(err)
	}
}
