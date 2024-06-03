package database

import (
	"context"
	"database/sql"
	"exoplanet/config"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4/database"
	dStub "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bunotel"
)

const (
	DEFAULT_DBCONNS = 75
)

func GetConnector(c config.Config) *sql.DB {
	return sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(getConnectionString(c)), pgdriver.WithPassword(c.DBPASS), pgdriver.WithUser(c.DBUSER)))

}

// NewDB will get a reusable DB provider with observability already setup.
func NewDB(c config.Config) bun.IDB {
	sqldb := GetConnector(c)

	db := bun.NewDB(sqldb, pgdialect.New())

	dbConns := getDbConns()
	db.SetMaxOpenConns(dbConns)
	db.SetMaxIdleConns(dbConns)

	db.AddQueryHook(bunotel.NewQueryHook(bunotel.WithDBName(c.DBNAME)))
	return db
}

func NewDBInstance(ctx context.Context, c config.Config) database.Driver {
	sqldb := GetConnector(c)
	instance, err := dStub.WithInstance(sqldb, &dStub.Config{})
	if err != nil {
		log.Println("Error: Could not connect to db", err)

	}
	return instance
}

func getConnectionString(c config.Config) string {
	return fmt.Sprintf("postgres://%s/%s?sslmode=disable", c.DBHOST, c.DBNAME)
}

func getDbConns() int {
	dbConns := DEFAULT_DBCONNS
	if os.Getenv("MAX_DB_CONNS") != "" {
		env, err := strconv.Atoi(os.Getenv("MAX_DB_CONNS"))
		if err == nil {
			dbConns = env
		}
	}
	return dbConns
}
