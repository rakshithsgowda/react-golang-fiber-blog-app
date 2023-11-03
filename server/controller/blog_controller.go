package controller

import "github.com/gofiber/fiber/v2"

// blog list page
// blog read/deatils page
// create/add blog into database
// update blog
// delete blog

func BlogList(c *fiber.Ctx)error{
	return c.JSON(fiber.Map{"message":"Welcome to blog web application"})
}

// func BlogDetail(){}

// func BlogCreate(){}

// func BlogUpdate(){}

// func BlogDelete(){}