package tests

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("ENV_NAME", "TEST")
	code := m.Run()
	// log.Println("RUNNING TEAR DOWN")
	// test_utils.RemoveTestDb()
	os.Setenv("ENV_NAME", "")
	os.Exit(code)
}
