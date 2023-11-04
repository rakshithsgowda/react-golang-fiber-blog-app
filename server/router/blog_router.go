package router

import (
	"blog/controller"

	"github.com/gofiber/fiber/v2"
)

// probable routes

// create blog -post method
// delete blog -Delte method
// update blog  -put method
// blog read/deatils - get meythod
// blog lits  - get method

func Routes_Setup(app *fiber.App){
	app.Get("/",controller.BlogList)
	app.Post("/",controller.BlogCreate)
	app.Put("/:id",controller.BlogUpdate)
	app.Delete("/:id",controller.BlogDelete)
	// app.Get("/:id",controller.BlogDetail)
}