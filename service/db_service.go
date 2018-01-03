package service

import (
	"database/sql"

	"github.com/adrenalinchik/gowebapp/repository"
)

func InitDB(dataSourceName string, dbName string, settings string) {
	repository.InitDB(dataSourceName, dbName, settings)
}

func GetDbConnection() *sql.DB {
	return repository.GetDbConn()
}

func CloseDbConn() {
	repository.GetDbConn().Close()
}

func GenerateDbData() {
	repository.GenerateDbData()
}

func DropDb(dbname string) {
	repository.DropDb(dbname)
}
