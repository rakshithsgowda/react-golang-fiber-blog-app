package main

import (
	"blog/database"
	"blog/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

// initialize
func init(){
	if err:=godotenv.Load(".env");err!=nil{
		log.Fatal("Error in loading .env file")
	}
	database.ConnectDB()
}
// connect to db
// log errors
// listen on port

func main() {

	// database.ConnectDB()
	mysqlDB,err:=database.DBConnection.DB()
	if err!=nil{
		panic("Error in mysql connection")
	}

	defer mysqlDB.Close()

	app := fiber.New()
	app.Use(logger.New())

	router.Routes_Setup(app)
	app.Listen(":8000")
}
