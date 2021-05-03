package tests

import (
	"log"
	"os"
	"simple-bank/api"
	"simple-bank/ent"
	"simple-bank/util"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var testDb *ent.Client
var config util.Config
var server *api.Server

func TestMain(m *testing.M) {
	env, err := util.LoadConfig("../", ".test")
	if err != nil {
		log.Fatalln("Env loading error ", err)
	}
	config = env
	testDb = SetupTestDb(&testing.T{})
	gin.SetMode(gin.TestMode)
	server = api.NewServer(testDb)

	defer testDb.Close()
	os.Exit(m.Run())
}
