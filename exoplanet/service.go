package exoplanet

import (
	"context"
)

func (e *ExoplanetScope) addExoplanet(ctx context.Context, planet *Exoplanet) error {
	err := e.addExoplanetquery(ctx, planet)
	return err
}
func (e *ExoplanetScope) listExoplanets(ctx context.Context) ([]Exoplanet, error) {
	var ex []Exoplanet
	ex, err := e.listExoplanetsquery(ctx)
	return ex, err
}

func (e *ExoplanetScope) getExoplanetByID(ctx context.Context, id int) (Exoplanet, error) {
	ex, err := e.getExoplanetByIDquery(ctx, id)
	return ex, err
}

func (e *ExoplanetScope) updateExoplanet(ctx context.Context, planet *Exoplanet) error {
	err := e.updateExoplanetquery(ctx, planet)
	return err
}

func (e *ExoplanetScope) deleteExoplanet(ctx context.Context, id int) error {
	err := e.deleteExoplanetquery(ctx, id)
	return err
}

func (e *ExoplanetScope) fuelEstimation(ctx context.Context, id int, crewCapacity int) (float64, error) {

	exoplanet, err := e.getExoplanetByIDquery(ctx, id)
	if err != nil {
		return 0, err
	}

	gravity := calculateGravity(exoplanet)
	fuelEstimation := calculateFuelEstimation(exoplanet.Distance, gravity, crewCapacity)
	return fuelEstimation, err
}

func calculateGravity(ex Exoplanet) float64 {
	if ex.Type == "GasGiant" {
		return 0.5 / (ex.Radius * ex.Radius)
	}
	return ex.Mass / (ex.Radius * ex.Radius)
}

func calculateFuelEstimation(distance float64, gravity float64, crewCapacity int) float64 {
	return distance / (gravity * gravity) * float64(crewCapacity)
}
