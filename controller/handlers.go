package controller

import (
	"net/http"

	"github.com/adrenalinchik/gowebapp/log"
)

func InitControllers() {
	// Owner handlers
	http.HandleFunc("/owner/get", getOwnerById)
	http.HandleFunc("/owner/getByEmail", getOwnerByEmail)
	http.HandleFunc("/owner/getByState", getByState)
	http.HandleFunc("/owner/getAll", getAll)
	http.HandleFunc("/owner/create", createOwner)
	http.HandleFunc("/owner/update", updateOwner)
	http.HandleFunc("/owner/delete", deleteOwner)

	// Vehicle handlers
	http.HandleFunc("/vehicle/create", createVehicle)
	http.HandleFunc("/vehicle/update", updateVehicle)
	http.HandleFunc("/vehicle/getByOwner", getVehiclesByOwner)
	log.Error.Fatal(http.ListenAndServe(":8080", nil))
}


//
//type badRequest struct{ error }
//
//type methodNotAllowed struct{ error }
//
//type notFound struct{ error }
//
//type unprocessableEntity struct{ error }
//
//func errorHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		err := f(w, r)
//		if err != nil {
//			switch err.(type) {
//			case methodNotAllowed:
//				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
//			case badRequest:
//				http.Error(w, err.Error(), http.StatusBadRequest)
//			case notFound:
//				http.Error(w, err.Error(), http.StatusNotFound)
//			case unprocessableEntity:
//				http.Error(w, err.Error(), http.StatusUnprocessableEntity)
//			default:
//				http.Error(w, err.Error(), http.StatusInternalServerError)
//				log.Error.Printf("handling %q: %v", r.RequestURI, err)
//			}
//		}
//	}
//}