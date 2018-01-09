package controller

import (
	"net/http"
	"encoding/json"
	"github.com/adrenalinchik/gowebapp/log"
	"github.com/adrenalinchik/gowebapp/service"
	"github.com/adrenalinchik/gowebapp/model"
)

var vehicleService = service.InitVehicleService()

func vehicleHandlers() {
	http.HandleFunc("/vehicle/create", CreateVehicle)
	http.HandleFunc("/vehicle/update", UpdateVehicle)
	http.HandleFunc("/vehicle/by-owner", GetVehiclesByOwner)
}

// GetVehiclesByOwner handles GET request to /vehicle/by-owner?owner_id={ownerId}.
// It returns a JSON encoded list of owner vehicles.
//
// Examples:
//
// 	req: GET /vehicle/by-owner?owner_id=1
// 	res: 200 [
// 				{"id":2,"owner_id":1,"model":"toyota","number":"112233","type":"hybrid","state":"active"},
// 				{"id":3,"owner_id":1,"model":"audi","number":"445566","type":"diesel","state":"active"}
// 			 ]
//
//	req: GET /vehicle/by-owner?owner_id=88
//	res: 404 no owner found with id = 88
func GetVehiclesByOwner(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id, err := parseID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Error.Println(err)
		return
	}
	vehicles, err := vehicleService.GetByOwner(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Error.Println(err)
		return
	}
	json.NewEncoder(w).Encode(vehicles)
}

// CreateVehicle handles POST request on /vehicle/create.
// The request body must contain JSON object with owner_id, model, number, type, state fields
// Number field have to be unique.
//
// Examples:
//
// 	req: POST /vehicle/create {"owner_id":2,"model":"diesel","number":"778899","type":"bmw","state":"active"}
// 	res: 200 {id:3,"owner_id":2,"model":"diesel","number":"778899","type":"bmw","state":"active"}
//
//	req: POST /vehicle/create {"owner_id":2,"model":"diesel","number":"778899","type":"bmw","state":"active"}
//	res: 422 vehicle with such number is already in the system
func CreateVehicle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var vehicle *model.Vehicle
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		log.Error.Println(err)
		return
	}
	if vehicle.Valid() {
		vehicle, err = vehicleService.Create(vehicle)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			log.Error.Println(err)
			return
		}
		json.NewEncoder(w).Encode(vehicle)
	} else {
		http.Error(w, "invalid vehicle body", http.StatusUnprocessableEntity)
		log.Error.Println("invalid vehicle body")
		return
	}
}

// UpdateVehicle handles PUT request on /vehicle/update.
// The request body must contain JSON encoded vehicle.
// Number field have to be unique.
//
// Examples:
//
// 	req: PUT /vehicle/update {id:1,"owner_id":2,"model":"diesel","number":"778899","type":"bmw","state":"active"}
// 	res: 200 {id:1,"owner_id":3,"model":"electro","number":"778899","type":"audi","state":"active"}
//
//	req: PUT /owner/update {id:2,"owner_id":3,"model":"diesel","number":"778899","type":"opel","state":"active"}
//	res: 422 vehicle with such number is already in the system
func UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var vehicle *model.Vehicle
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		log.Error.Println(err)
		return
	}
	vehicle, err = vehicleService.Update(vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		log.Error.Println(err)
		return
	}
	json.NewEncoder(w).Encode(vehicle)
}
