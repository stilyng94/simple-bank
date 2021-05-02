package tests

import (
	"log"
	"simple-bank/ent"
	"simple-bank/ent/enttest"
	"simple-bank/ent/migrate"
	"simple-bank/operation"
	"testing"
)

func SetupTestDb(t *testing.T) *ent.Client {
	drv := operation.SetupDbDriver("postgresql://root:secret@127.0.0.1:5432/test_bank")

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(log.Println), ent.Driver(drv)),
		enttest.WithMigrateOptions(migrate.WithDropColumn(true), migrate.WithDropIndex(true)),
	}

	client := enttest.NewClient(t, opts...)
	return client

}
