package tests

import (
	"embed"
	"go-payments/configs/database"
	"log"
	"os"

	"github.com/pressly/goose/v3"
)

//go:embed test_migrations/*.sql
var EmbedTestMigrations embed.FS

func InitTestDb() {
	os.Create("test.db")
	log.Print("test database created")
	db := database.GetDb()
	if db == nil {
		log.Fatalf("database couldn't be accessed")
	}
	defer db.Close()
	goose.SetBaseFS(EmbedTestMigrations)
	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatalf("db migrations couldn't be executed in set dialect phase | %v", err)
	}
	if err := goose.Up(db, "test_migrations"); err != nil {
		log.Fatalf("db migrations couldn't be executed with embed migrations %+v | %v", EmbedTestMigrations, err)
	}
}

func RemoveTestDb() {
	err := os.Remove("test.db")
	if err != nil {
		log.Printf("test database remove error | %v", err)
		return
	}
	log.Print("test database removed")
}
