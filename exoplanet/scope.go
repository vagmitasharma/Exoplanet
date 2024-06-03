package exoplanet

import (
	"exoplanet/config"

	"github.com/uptrace/bun"
)

type ExoplanetScope struct {
	db     bun.IDB
	config config.Config
}

// NewScope will setup all dependencies for this feature (use dependency injection).
func NewScope(db bun.IDB, c config.Config) *ExoplanetScope {
	s := &ExoplanetScope{
		db:     db,
		config: c,
	}
	return s
}
