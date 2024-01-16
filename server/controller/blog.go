package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marco210/blog-app/database"
	"github.com/marco210/blog-app/model"
)

func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"msg": "BlogList",
	}

	db := database.DBConn

	var records []model.Blog

	db.Find(&records)
	context["blog_records"] = records

	c.Status(200)
	return c.JSON(context)
}

func BlogCreate(c *fiber.Ctx) error {
	

	db := database.DBConn
	var records []model.Blog
	db.Create(&records)

	context := fiber.Map{
		"msg": "Added a Blog ",
	}

	c.Status(200)
	return c.JSON(context)
	
}

func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"msg": "Update a Blog",
	}
	c.Status(200)
	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"msg": "Delete a blog",
	}
	c.Status(200)
	return c.JSON(context)
}
