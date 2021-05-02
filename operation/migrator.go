package operation

import (
	"context"
	"log"
	"simple-bank/ent"
	"simple-bank/ent/migrate"
)

func MigrateDb(dbClient *ent.Client) {
	ctx := context.Background()
	if err := dbClient.Schema.Create(ctx, migrate.WithDropColumn(true), migrate.WithDropIndex(true)); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
}
