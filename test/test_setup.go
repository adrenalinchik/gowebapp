package test

import (
	"io/ioutil"
	"os"

	"github.com/adrenalinchik/gowebapp/log"
	"github.com/adrenalinchik/gowebapp/service"
)

//var testDataSourceName = "root:@tcp(127.0.0.1:3306)/"
var testDataSourceName = "tarasfihurnyak:12345678@tcp(parkingdbinstance.cdn09iyfeibg.us-east-1.rds.amazonaws.com:3306)/"
var testDbName = "goparkingtests"
var settings = "?parseTime=true"


func Setup() {
	log.InitLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	service.InitDB(testDataSourceName, testDbName, settings)
	service.GenerateDbData()
}

func Shutdown() {
	service.DropDb(testDbName)
	service.CloseDbConn()
}
