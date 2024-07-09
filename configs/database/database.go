package database

import (
	"database/sql"
	"embed"
	"log"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func GetTestDb() *sql.DB {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetDb() *sql.DB {
	env_name := os.Getenv("ENV_NAME")
	if env_name == "TEST" {
		return GetTestDb()
	}
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
