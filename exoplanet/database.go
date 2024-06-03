package exoplanet

import (
	"context"
	"log"
)

func (e *ExoplanetScope) addExoplanetquery(ctx context.Context, planet *Exoplanet) error {

	_, err := e.db.NewInsert().Model(&planet).Exec(ctx)
	if err != nil {
		log.Println("Error: " + err.Error())
		return err
	}
	return err
}
func (e *ExoplanetScope) listExoplanetsquery(ctx context.Context) ([]Exoplanet, error) {
	var exoplanets []Exoplanet
	err := e.db.NewSelect().Model(&exoplanets).Scan(ctx)
	if err != nil {
		log.Println("Error: " + err.Error())
		return exoplanets, err
	}

	return exoplanets, err
}

func (e *ExoplanetScope) getExoplanetByIDquery(ctx context.Context, id int) (Exoplanet, error) {

	var exoplanet Exoplanet
	err := e.db.NewSelect().Model(&exoplanet).Where("id = ?", id).Scan(ctx)
	if err != nil {
		log.Println("Error: " + "Exoplanet not found")
		return exoplanet, err
	}
	return exoplanet, err
}

func (e *ExoplanetScope) updateExoplanetquery(ctx context.Context, planet *Exoplanet) error {

	var exoplanet Exoplanet
	_, err := e.db.NewUpdate().Model(&exoplanet).Where("id = ?", &planet.ID).Exec(ctx)
	if err != nil {
		log.Println("Error: " + err.Error())
		return err
	}
	return err
}

func (e *ExoplanetScope) deleteExoplanetquery(ctx context.Context, id int) error {

	_, err := e.db.NewDelete().Model(&Exoplanet{}).Where("id = ?", id).Exec(ctx)
	if err != nil {
		log.Println("Error: " + err.Error())
		return err
	}
	return err
}
