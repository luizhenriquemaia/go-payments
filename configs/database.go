package configs

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Get_db() *sql.DB {
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

func Init_db() {
	db := Get_db()
	repository := New_repository(db)

	if err := repository.Migrate(); err != nil {
		log.Fatal(err)
	}
}
