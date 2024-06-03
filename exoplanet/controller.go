package exoplanet

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (e *ExoplanetScope) Route(base *mux.Router, api *mux.Router) {
	r := api.PathPrefix("/").Subrouter()
	r.HandleFunc("/", e.AddExoplanet).Methods("POST")
	r.HandleFunc("/", e.ListExoplanets).Methods("GET")
	r.HandleFunc("/{id}", e.GetExoplanetByID).Methods("GET")
	r.HandleFunc("/{id}", e.UpdateExoplanet).Methods("PUT")
	r.HandleFunc("/{id}", e.DeleteExoplanet).Methods("DELETE")
	r.HandleFunc("/{id}/fuel_estimation", e.FuelEstimation).Methods("GET")

}
func (e *ExoplanetScope) AddExoplanet(w http.ResponseWriter, r *http.Request) {

	var exoplanet Exoplanet
	err := json.NewDecoder(r.Body).Decode(&exoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = e.addExoplanet(r.Context(), &exoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func (e *ExoplanetScope) ListExoplanets(w http.ResponseWriter, r *http.Request) {
	var exoplanets []Exoplanet
	exoplanets, err := e.listExoplanets(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(exoplanets)
}

func (e *ExoplanetScope) GetExoplanetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid exoplanet ID", http.StatusBadRequest)
		return
	}

	var exoplanet Exoplanet
	exoplanet, err = e.getExoplanetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(exoplanet)
}

func (e *ExoplanetScope) UpdateExoplanet(w http.ResponseWriter, r *http.Request) {

	var exoplanet Exoplanet
	err := json.NewDecoder(r.Body).Decode(&exoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = e.updateExoplanet(r.Context(), &exoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (e *ExoplanetScope) DeleteExoplanet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid exoplanet ID", http.StatusBadRequest)
		return
	}

	err = e.deleteExoplanet(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (e *ExoplanetScope) FuelEstimation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid exoplanet ID", http.StatusBadRequest)
		return
	}

	crewCapacityStr := r.URL.Query().Get("crew_capacity")
	if crewCapacityStr == "" {
		http.Error(w, "Crew capacity not provided", http.StatusBadRequest)
		return
	}

	crewCapacity, err := strconv.Atoi(crewCapacityStr)
	if err != nil {
		http.Error(w, "Invalid crew capacity", http.StatusBadRequest)
		return
	}
	fuelEstimation, err := e.fuelEstimation(r.Context(), id, crewCapacity)
	if err != nil {
		http.Error(w, "Could not calculate", http.StatusBadRequest)
		return
	}

	response := map[string]float64{"fuel_estimation": fuelEstimation}
	json.NewEncoder(w).Encode(response)
}
