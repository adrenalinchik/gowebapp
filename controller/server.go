package controller

import (
	"net/http"

	"github.com/adrenalinchik/gowebapp/log"
)

func RegisterHandlers() {
	ownerHandlers()
	vehicleHandlers()
	log.Error.Fatal(http.ListenAndServe(":8080", nil))
}
