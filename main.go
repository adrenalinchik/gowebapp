package main

import (
	"io/ioutil"
	"os"

	"github.com/adrenalinchik/gowebapp/controller"
	"github.com/adrenalinchik/gowebapp/service"
	"github.com/adrenalinchik/gowebapp/log"
)

func main() {
	log.InitLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	log.Info.Println("Starting...")
	//v := []model.Vehicle{
	//	model.Vehicle{1},
	//	model.Vehicle{2},
	//}
	//owner := &model.Owner{
	//	64,
	//	"Taras",
	//	"Fihurnyak",
	//	model.Male,
	//	time.Now(),
	//	model.Active,
	//	v}

	service.InitDB("root:@tcp(127.0.0.1:3306)/", "test1", "?parseTime=true")
	log.Info.Println("Migration finished successfully")
	controller.InitControllers()

	defer service.CloseDbConn()
}
