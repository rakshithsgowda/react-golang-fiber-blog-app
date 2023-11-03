package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// connect to the database (mysql)

// requird parameters -  host, user, password nmae of db

var DBConnection *gorm.DB

func ConnectDB() {
	
	dsn :=  "root:admin@tcp(127.0.0.1:3306)/fiber_blog_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic("Database connection failed.")
	}
	log.Println("Connection successful")

	DBConnection = db
}
