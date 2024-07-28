package main

import (
	"embed"
	"go-payments/configs/database"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

func getWorkDir() string {
	work_dir, err := os.Getwd()
	if err != nil {
		log.Fatal("working directory wasn't setted")
	}
	_, err = os.Stat(work_dir + "/temp/tests/")
	if os.IsNotExist(err) {
		err = os.Mkdir(work_dir+"/temp/tests", 0777)
		if err != nil {
			log.Fatalf("couldn't create temp test dir: %v", err)
		}
	}
	return work_dir + "/temp/tests/"
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

func copyFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("couldn't open file to copy: %v", path)
	}
	defer file.Close()
	new_file_name := strings.ReplaceAll(filepath.ToSlash(path), "/", "_")
	dest_file, err := os.Create("temp/tests/" + new_file_name)
	if err != nil {
		log.Fatalf("couldn't create copy test file: %v", path)
	}
	defer dest_file.Close()
	_, err = io.Copy(dest_file, file)
	if err != nil {
		log.Fatalf("couldn't copy test file: %v", path)
	}
}

func copyTestFiles() {
	log.Print("copying test files")
	name_regex, err := regexp.Compile("^.+_test.go")
	if err != nil {
		log.Fatal("fail to parse test filename regex", err)
	}
	number_founded_files := 0
	err = filepath.Walk("./internal", func(path string, info os.FileInfo, err error) error {
		if name_regex.MatchString(info.Name()) {
			number_founded_files += 1
			copyFile(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal("fail to walk directory ", err)
	}
	log.Println("founded files", number_founded_files)
}

func removeTestFiles() {
	dir := "temp/tests/"
	os.RemoveAll(dir)
	os.Mkdir(dir, 0777)
}

func (suite *DefaultTestSuite) SetupSuite() {
	os.Setenv("ENV_NAME", "TEST")
	main_dir := getWorkDir()
	os.Setenv("MAIN_DIR", main_dir)
	initTestDb()
	copyTestFiles()
}

func (suite *DefaultTestSuite) TearDownSuite() {
	err := os.Remove("test.db")
	if err != nil {
		log.Printf("test database remove error | %v", err)
	} else {
		log.Print("test database removed")
	}
	removeTestFiles()
}
