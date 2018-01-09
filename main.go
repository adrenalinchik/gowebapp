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
	log.Info.Println("starting...")
	service.InitDB("root:@tcp(127.0.0.1:3306)/", "test1", "?parseTime=true")
	defer service.CloseDbConn()
	log.Info.Println("migration finished successfully")
	controller.RegisterHandlers()
}
