package operation

import (
	"simple-bank/ent"
)

// Connect to database
func DbConn(databaseUrl string) *ent.Client {
	drv := SetupDbDriver(databaseUrl)
	// Create an ent.Driver from `db`.
	return ent.NewClient(ent.Driver(drv))
}
