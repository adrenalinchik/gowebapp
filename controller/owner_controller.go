package controller

import (
	"net/http"
	"encoding/json"
	"github.com/adrenalinchik/gowebapp/service"
	"github.com/adrenalinchik/gowebapp/model"
)

var ownerService = service.InitOwnerService()

func ownerHandlers() {
	http.HandleFunc("/owner", GetOwner)
	http.HandleFunc("/owner/by-email", GetOwnerByEmail)
	http.HandleFunc("/owner/by-state", GetByState)
	http.HandleFunc("/owner/all", GetAll)
	http.HandleFunc("/owner/create", CreateOwner)
	http.HandleFunc("/owner/update", UpdateOwner)
	http.HandleFunc("/owner/delete", DeleteOwner)
}

// GetOwner handles GET request to /owner?id={ownerId}.
// It returns a JSON encoded owner.
//
// Examples:
//
// 	req: GET /owner?id=1
// 	res: 200 {"id":1,"firstname":"TestName","lastname":"TestLastName","gender":"male","dob":"1988-01-02T00:00:00Z","email":"test@gmail.com","state":"active","vehicles":[]}
//
//	req: GET /owner?id=67
//	res: 404 no owner found with id = 67
func GetOwner(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id, err := parseID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	owner, err := ownerService.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(owner)
}

// GetOwnerByEmail handles GET request to /owner/by-email?email={ownerEmail}.
// It returns a JSON encoded owner.
//
// Examples:
//
// 	req: GET /owner/by-email?email=test@gmail.com
// 	res: 200 {"id":1,"firstname":"TestName","lastname":"TestLastName","gender":"male","dob":"1988-01-02T00:00:00Z","email":"test@gmail.com","state":"active","vehicles":[]}
//
//	req: GET /owner/by-email?email=incorrect@gmail.com
//	res: 404 no owner found with email = incorrect@gmail.com
func GetOwnerByEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "email field is empty", http.StatusBadRequest)
		return

	}
	owner, err := ownerService.GetByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(owner)
}

// GetAll handles GET request to /owner/all.
// There's no parameters and it returns a JSON encoded list of owners.
//
// Example:
//
// 	req: GET /owner/all
//  res: 200[
// 			  {"id":1,"firstname":"TestName","lastname":"TestLastName","gender":"male","dob":"1988-01-02T00:00:00Z","email":"test@gmail.com","state":"active","vehicles":[]},
//			  {"id":2,"firstname":"TestName2","lastname":"TestLastName2","gender":"male","dob":"1988-01-02T00:00:00Z","email":"test2@gmail.com","state":"inactive","vehicles":[]},
//			  {"id":3,"firstname":"TestName3","lastname":"TestLastName3","gender":"female","dob":"1988-01-02T00:00:00Z","email":"test3@gmail.com","state":"inactive","vehicles":[]}
//          ]
func GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	owners, err := ownerService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(owners)
}

// GetByState handles GET request to /owner/by-state?state={ownerState}.
// It returns a JSON encoded list of owners by it state. State have to be active or inactive.
//
// Examples:
//
// 	req: GET /owner/by-state?state=inactive
// 	res: 200[
//			  {"id":2,"firstname":"TestName2","lastname":"TestLastName2","gender":"male","dob":"1988-01-02T00:00:00Z","email":"test2@gmail.com","state":"inactive","vehicles":[]},
//			  {"id":3,"firstname":"TestName3","lastname":"TestLastName3","gender":"female","dob":"1988-01-02T00:00:00Z","email":"test3@gmail.com","state":"inactive","vehicles":[]}
//          ]
//
//	req: GET /owner/by-state?state=somestate
//	res: 400 state field is incorrect
func GetByState(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	state := r.FormValue("state")
	if state == "" {
		http.Error(w, "state field is empty", http.StatusBadRequest)
		return
	}
	var owners [] *model.Owner
	var err error

	switch state {
	case "active":
		owners, err = ownerService.GetByState(model.ACTIVE)
	case "inactive":
		owners, err = ownerService.GetByState(model.INACTIVE)
	default:
		http.Error(w, "state field is incorrect", http.StatusBadRequest)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(owners)
}

// CreateOwner handles POST request on /owner/create.
// The request body must contain JSON object with firstname, lastname, gender, dob, email, state fields
// Email field have to be unique.
// Vehicles field is optional and it could contain the list of owner vehicles.
//
// Examples:
//
// 	req: POST /owner/create {"firstname": "Test", "lastname": "TestTest", "gender": "male", "dob": "1988-01-02T15:04:05Z", "email": "test1@gmail.com", "state": "inactive"}
// 	res: 200 {"id":1,"firstname":"Test","lastname":"TestTest","gender":"male","dob":"1988-01-02T00:00:00Z","email":"test1@gmail.com","state":"inactive","vehicles":[]}
//
//	req: POST /owner/create {"firstname": "Test", "lastname": "TestTest", "gender": "male", "dob": "1988-01-02T15:04:05Z", "email": "test1@gmail.com", "state": "inactive"}
//	res: 422 owner with such email is already in the system
func CreateOwner(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var owner *model.Owner
	err := json.NewDecoder(r.Body).Decode(&owner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	owner, err = ownerService.Create(owner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	json.NewEncoder(w).Encode(owner)
}

// UpdateOwner handles PUT request on /owner/update.
// The request body must contain JSON encoded owner.
// Email field have to be unique.
//
// Examples:
//
// 	req: PUT /owner/update?id=1 {"Id":1,"firstname": "Test", "lastname": "TestTest", "gender": "male", "dob": "1988-01-02T15:04:05Z", "email": "test@gmail.com", "state": "active"}
// 	res: 200 {"Id":1,"firstname": "TestUpdated", "lastname": "TestTestUpdated", "gender": "female", "dob": "1988-01-02T15:04:05Z", "email": "testupdated@gmail.com", "state": "active", "vehicles":[]}
//
//	req: PUT /owner/update {"Id":2,"firstname": "FirstName", "lastname": "LastName", "gender": "male", "dob": "1988-01-02T15:04:05Z", "email": "testupdated@gmail.com", "state": "active"}
//	res: 422 owner with such email is already in the system
func UpdateOwner(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var owner *model.Owner
	err := json.NewDecoder(r.Body).Decode(&owner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	owner, err = ownerService.Update(owner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	json.NewEncoder(w).Encode(owner)
}

// DeleteOwner handles DELETE request to /owner/delete?id={ownerId}.
//
// Examples:
//
// 	req: GET /owner/delete?id=1
// 	res: 200
//
//	req: GET /owner/delete?id=88
//	res: 404 no owner found with id = 88
func DeleteOwner(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id, err := parseID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = ownerService.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}
