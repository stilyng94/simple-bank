package test

import (
	"log"
	"simple-bank/ent"
	"simple-bank/ent/enttest"
	"simple-bank/ent/migrate"
	"simple-bank/operation"
	"testing"
)

func SetupTestDb(t *testing.T) *ent.Client {
	drv := operation.SetupDbDriver(config.DB_SOURCE)

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(log.Println), ent.Driver(drv)),
		enttest.WithMigrateOptions(migrate.WithDropColumn(true), migrate.WithDropIndex(true)),
	}

	client := enttest.NewClient(t, opts...)
	return client

}
