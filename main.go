package main

import (
	"simple-bank/api"
	"simple-bank/operation"
)

func main() {
	dbClient := operation.DbConn("postgresql://root:secret@127.0.0.1:5432/simple_bank")

	operation.MigrateDb(dbClient)

	server := api.NewServer(dbClient)

	err := server.Start("0.0.0.0:5000")
	if err != nil {
		panic(err)
	}

	defer dbClient.Close()

}
