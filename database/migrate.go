package database

import (
	"context"
	"exoplanet/config"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrate will automatically update the DB to the latest schema.
func Migrate(ctx context.Context, c config.Config) {
	enabled := os.Getenv("MIGRATE_DB")
	if !(enabled == "" || enabled == "true") {
		log.Println("Info: " + "Migrate DB is disabled")
		return
	}
	log.Println("Info: " + "Migrate DB to latest schema")
	migrationFilePath := c.MIGRATIONSCRIPTPATH

	m, err := migrate.NewWithDatabaseInstance(migrationFilePath, c.DBNAME, NewDBInstance(ctx, c))
	if err != nil {
		log.Fatal("Migrate DB failed to init", err)
	}

	err = m.Up()
	if err != nil && err.Error() == "no change" {
		log.Println("Info: " + "Migrate DB detected no change")
	} else if err != nil {
		log.Fatal("Migrate DB failed", err)
	} else {
		log.Println("Info: " + "Migrate DB updated schema successfully")
	}
}
