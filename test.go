package main

import (
	"embed"
	"go-payments/configs/database"
	"log"
	"os"
	"testing"

	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/suite"
)

//go:embed configs/database/test_migrations/*.sql
var embed_test_migrations embed.FS

type DefaultTestSuite struct {
	suite.Suite
}

func TestTestApisSuite(test *testing.T) {
	test.Setenv("ENV_NAME", "TEST")
	suite.Run(test, &DefaultTestSuite{})
}

func getMainDir() string {
	work_dir, err := os.Getwd()
	if err != nil {
		log.Fatal("working directory wasn't setted")
	}
	return work_dir
}

func initTestDb() {
	os.Create("test.db")
	log.Print("test database created")
	db := database.GetDb()
	defer db.Close()
	log.Printf("test database founded at: %v | %v", db, db.Stats())
	if db == nil {
		log.Fatal("database couldn't be accessed")
	}
	goose.SetBaseFS(embed_test_migrations)
	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatalf("db migrations couldn't be executed in set dialect phase | %v", err)
	}
	if err := goose.Up(db, "configs/database/test_migrations"); err != nil {
		log.Fatalf("db migrations couldn't be executed | %v", err)
	}
}

func (suite *DefaultTestSuite) SetupSuite() {
	os.Setenv("ENV_NAME", "TEST")
	main_dir := getMainDir()
	os.Setenv("MAIN_DIR", main_dir)
	initTestDb()
}

func (suite *DefaultTestSuite) TearDownSuite() {
	err := os.Remove("test.db")
	if err != nil {
		log.Printf("test database remove error | %v", err)
	} else {
		log.Print("test database removed")
	}
}
