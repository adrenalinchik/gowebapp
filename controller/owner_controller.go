package controller

import (
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/adrenalinchik/gowebapp/service"
	"github.com/adrenalinchik/gowebapp/model"
)

var ownerService = service.InitOwnerService()

func getOwnerById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "id field is empty", http.StatusBadRequest)
		return
	}
	ownerid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "not a number: " + id, http.StatusBadRequest)
		return
	}
	owner, err := ownerService.Get(ownerid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(owner)
}

func getOwnerByEmail(w http.ResponseWriter, r *http.Request)  {
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

func getAll(w http.ResponseWriter, r *http.Request) {
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

func getByState(w http.ResponseWriter, r *http.Request) {
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
	if state == "inactive" {
		owners, err = ownerService.GetByState(model.INACTIVE)
	} else {
		owners, err = ownerService.GetByState(model.ACTIVE)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(owners)
}

func createOwner(w http.ResponseWriter, r *http.Request) {
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

func updateOwner(w http.ResponseWriter, r *http.Request) {
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

func deleteOwner(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "id field is empty", http.StatusBadRequest)
		return
	}
	ownerid, _ := strconv.ParseInt(id, 10, 64)
	err := ownerService.Delete(ownerid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}
