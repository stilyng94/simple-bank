package operation

import (
	"simple-bank/ent"
)

// Connect to database
func DbConn(databaseUrl string) *ent.Client {
	drv := SetupDbDriver(databaseUrl)
	return ent.NewClient(ent.Driver(drv))
}
