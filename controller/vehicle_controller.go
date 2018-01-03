package controller

import (
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/adrenalinchik/gowebapp/log"
	"github.com/adrenalinchik/gowebapp/service"
	"github.com/adrenalinchik/gowebapp/model"
)

var vehicleService = service.InitVehicleService()

func getVehiclesByOwner(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	ownerId := r.FormValue("owner_id")
	if ownerId == "" {
		http.Error(w, http.StatusText(400), 400)
		log.Error.Println("OwnerId field is empty.")
		return
	}
	id, _ := strconv.ParseInt(ownerId, 10, 64)
	vehicles, err := vehicleService.GetByOwner(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		log.Error.Println(err)
		return
	}
	json.NewEncoder(w).Encode(vehicles)
}

func createVehicle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	var vehicle *model.Vehicle
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		http.Error(w, err.Error(), 422)
		log.Error.Println(err)
		return
	}
	if vehicle.Valid() {
		vehicle, err = vehicleService.Create(vehicle)
		if err != nil {
			http.Error(w, err.Error(), 422)
			log.Error.Println(err)
			return
		}
		json.NewEncoder(w).Encode(vehicle)
	} else {
		http.Error(w, "Invalid vehicle body", 422)
		log.Error.Println("Invalid vehicle body")
		return
	}
}

func updateVehicle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	var vehicle *model.Vehicle
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		http.Error(w, err.Error(), 422)
		log.Error.Println(err)
		return
	}
	vehicle, err = vehicleService.Update(vehicle)
	if err != nil {
		http.Error(w, err.Error(), 422)
		log.Error.Println(err)
		return
	}
	json.NewEncoder(w).Encode(vehicle)
}
