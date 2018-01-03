package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/adrenalinchik/gowebapp/log"
	"github.com/adrenalinchik/gowebapp/model"
	"time"
	"strconv"
)

var db *sql.DB

func InitDB(dataSourceName string, dbName string, settings string) {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Error.Panicln(err)
	}
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		log.Error.Panicln(err)
	}
	db.Close()
	db, err = sql.Open("mysql", dataSourceName+dbName+settings)
	if err != nil {
		log.Error.Panicln(err)
	}
	if err = db.Ping(); err != nil {
		log.Error.Panicln(err)
	} else {
		createTables()
	}
}

func GetDbConn() *sql.DB {
	return db
}

func GenerateDbData() {
	for i := 0; i < 10; i++ {
		num := strconv.Itoa(i)
		InitOwnerRepo().Insert(&model.Owner{
			Firstname: "Firstname" + num,
			Lastname:  "Lastname" + num,
			Gender:    model.MALE,
			Dob:       time.Now(),
			Email:     "test" + num + "@gmail.com",
			State:     model.ACTIVE,
		})
	}
}

func DropDb(dbname string) {
	_, err := db.Exec("DROP SCHEMA IF EXISTS " + dbname)
	if err != nil {
		log.Error.Println(err)
	}
}

func createTables() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS OWNER (ID BIGINT(20) PRIMARY KEY AUTO_INCREMENT, " +
		"FIRSTNAME VARCHAR (255), LASTNAME VARCHAR (255), GENDER VARCHAR (255), DOB DATE, EMAIL VARCHAR (255), STATE VARCHAR (255))")
	if err != nil {
		log.Error.Println(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS VEHICLE (ID BIGINT(20) PRIMARY KEY AUTO_INCREMENT, OWNER_ID BIGINT (20), " +
		"TYPE VARCHAR (255), NUMBER VARCHAR (255), MODEL VARCHAR (255), STATE VARCHAR (255))")
	if err != nil {
		log.Error.Println(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS GARAGE (ID BIGINT(20) PRIMARY KEY AUTO_INCREMENT, PARKING_ID BIGINT (20), " +
		"TYPE VARCHAR (255), SQUARE FLOAT)")
	if err != nil {
		log.Error.Println(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS PARKING (ID BIGINT(20) PRIMARY KEY AUTO_INCREMENT, ADDRESS VARCHAR (255)," +
		" GARAGES_NUMBER INT(11))")
	if err != nil {
		log.Error.Println(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS RESERVATION (ID BIGINT(20) PRIMARY KEY AUTO_INCREMENT,BEGIN DATETIME, " +
		"END DATETIME, GARAGE_ID BIGINT(20), OWNER_ID BIGINT(20), PARKING_ID BIGINT(20))")
	if err != nil {
		log.Error.Println(err)
	}
}
