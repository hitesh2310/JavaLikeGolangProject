package database

import (
	"fmt"
	"main/logs"
	"main/pkg/constants"
	"main/pkg/models"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var dbConn *gorm.DB

func EstablishDbConnection() {
	logs.InfoLog("Establishing DB connection")
	host := constants.ApplicationConfig.Database.Host
	port := constants.ApplicationConfig.Database.Port
	username := constants.ApplicationConfig.Database.Username
	password := constants.ApplicationConfig.Database.Password
	dbname := constants.ApplicationConfig.Database.Dbname

	logs.InfoLog(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname))

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", username, password, host, port, dbname)))
	if err != nil {
		logs.InfoLog(err.Error())
		logs.ErrorLog("Not connected %v", err)
	}
	db.AutoMigrate(models.Product{}, models.User{})
	dbConn = db
	logs.InfoLog("Connected to the database")

}

func GetDBConn() *gorm.DB {
	EstablishDbConnection()

	return dbConn
}

func GetDBConnQA() *gorm.DB {
	logs.InfoLog("Establishing DB connection")
	host := "localhost"
	port := "3306"
	username := "hiteshm"
	password := "qwerty"
	dbname := "ecomdb"

	logs.InfoLog(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname))

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", username, password, host, port, dbname)))
	if err != nil {
		logs.InfoLog(err.Error())
		logs.ErrorLog("Not connected %v", err)
	}
	db.AutoMigrate(models.Product{}, models.User{})
	dbConn = db
	logs.InfoLog("Connected to the database")

	return db
}
