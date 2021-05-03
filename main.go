package main

import (
	"log"
	"simple-bank/api"
	"simple-bank/operation"
	"simple-bank/util"
)

func main() {
	config, err := util.LoadConfig(".", ".dev")
	if err != nil {
		log.Fatalln("Env loading error ", err)
	}
	dbClient := operation.DbConn(config.DB_SOURCE)

	operation.MigrateDb(dbClient)

	server := api.NewServer(dbClient)

	err = server.Start(config.ServerAddress)
	if err != nil {
		panic(err)
	}

	defer dbClient.Close()

}
