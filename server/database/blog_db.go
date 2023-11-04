package database

import (
	"blog/model"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// connect to the database (mysql)

// requird parameters -  host, user, password nmae of db

var DBConnection *gorm.DB

func ConnectDB() {
	

user:=os.Getenv("db_user")
password:=os.Getenv("db_password")
dbname:=os.Getenv("db_name")
dsn:= user+":"+password+"@tcp(127.0.0.1:3306)/"+dbname+"?charset=utf8mb4&parseTime=True&loc=Local"
// dsn:=os.Getenv("DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic("Database connection failed.")
	}
	log.Println("Connection successful")
	db.AutoMigrate(new(model.Blog))

	DBConnection = db
}
