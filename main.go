package main

import (
	"context"
	"exoplanet/api"
	"exoplanet/config"
	"exoplanet/database"
	"exoplanet/exoplanet"
	"log"
)

func main() {

	c, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	database.Migrate(ctx, c)

	db := database.NewDB(c)
	e := exoplanet.NewScope(db, c)

	errChan := make(chan error)
	api.Serve(ctx, []api.ScopeRouter{e}, errChan, c)
	log.Fatal(ctx, "Fatal error", <-errChan)
}
