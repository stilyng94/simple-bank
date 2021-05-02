package main

import "simple-bank/operation"

func main() {
	dbClient := operation.DbConn("postgresql://root:secret@127.0.0.1:5432/simple_bank")

	operation.MigrateDb(dbClient)

	defer dbClient.Close()

}
