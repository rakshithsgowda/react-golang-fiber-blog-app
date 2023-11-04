package controller

import (
	"blog/database"
	"blog/model"

	"github.com/gofiber/fiber/v2"
)

// blog list page
func BlogList(c *fiber.Ctx)error{
	context:= fiber.Map{
		"statusText":"Ok",
		"msg":"Blog List",
	}

	db:= database.DBConnection
	var records []model.Blog

	db.Find(&records)

	context["blog_records"] = records
	c.Status(200)
	return c.JSON(context)
}

// create/add blog into database
func BlogCreate(c *fiber.Ctx)error{
	context:= fiber.Map{
		"statustext":"Ok",
		"msg":"Add a Blog",
	}
	c.Status(201)
	return c.JSON(context)
}


// update blog
func BlogUpdate(c *fiber.Ctx)error{
		context:=fiber.Map{
		"statusText":"Ok",
		"msg":"Blog Details updated",
		}
	c.Status(200)
	return c.JSON(context)
	}	

// delete blog
func BlogDelete(c *fiber.Ctx)error{
	context:= fiber.Map{
		"statusText":"Ok",
		"msg":"Delete Blog for the given ID",
	}
	c.Status(200)
	return c.JSON(context)
}


// blog read/deatils page
// func BlogDetail(){}
