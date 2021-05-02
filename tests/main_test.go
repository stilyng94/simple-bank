package tests

import (
	"os"
	"simple-bank/ent"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var testDb *ent.Client

func TestMain(m *testing.M) {
	testDb = SetupTestDb(&testing.T{})
	defer testDb.Close()
	os.Exit(m.Run())
}
