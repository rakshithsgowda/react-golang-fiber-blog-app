package main

import (
	"blog/database"
	"blog/router"

	"github.com/gofiber/fiber/v2"
)

// initialize

// connect to db
// log errors
// listen on port

func main() {

	database.ConnectDB()

	mysqlDB,err:=database.DBConnection.DB()
	if err!=nil{
		panic("Error in mysql connection")
	}

	defer mysqlDB.Close()

	app := fiber.New()

	router.Routes_Setup(app)
	app.Listen(":8000")
}
