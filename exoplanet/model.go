package exoplanet

import "github.com/uptrace/bun"

type ExoplanetType string

const (
	GasGiant    ExoplanetType = "GasGiant"
	Terrestrial ExoplanetType = "Terrestrial"
)

type Exoplanet struct {
	bun.BaseModel `bun:"exoplanets"`

	ID          int     `json:"id" bun:"id"`
	Name        string  `json:"name" bun:"name"`
	Description string  `json:"description" bun:"description"`
	Distance    float64 `json:"distance" bun:"distance"`
	Radius      float64 `json:"radius" bun:"radius"`
	Mass        float64 `json:"mass" bun:"mass"`
	Type        string  `json:"type" bun:"type"`
}
