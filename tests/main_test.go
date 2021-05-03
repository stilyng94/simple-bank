package tests

import (
	"log"
	"os"
	"simple-bank/ent"
	"simple-bank/util"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var testDb *ent.Client
var config util.Config

func TestMain(m *testing.M) {
	envv, err := util.LoadConfig("../", ".test")
	if err != nil {
		log.Fatalln("Env loading error ", err)
	}
	config = envv
	testDb = SetupTestDb(&testing.T{})
	defer testDb.Close()
	os.Exit(m.Run())
}
