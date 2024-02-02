package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/marco210/blog-app/database"
	"github.com/marco210/blog-app/model"
)

func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "BlogList",
		"statusText": "OK",
	}

	db := database.DBConn

	var records []model.Blog

	db.Find(&records)
	context["blog_records"] = records

	c.Status(200)
	return c.JSON(context)
}

// Blog detail Page
func BlogListById(c *fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "",
		"statusText": "",
	}

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not found")
		context["msg"] = "Record not found"

		c.Status(404)
		return c.JSON(context)
	}

	context["statusText"] = "OK"
	context["msg"] = "Blog Detail"
	context["blog_records"] = record

	c.Status(200)
	return c.JSON(context)
}

func BlogCreate(c *fiber.Ctx) error {

	context := fiber.Map{
		"msg":        "Added a Blog ",
		"statusText": "OK",
	}

	record := new(model.Blog)
	if err := c.BodyParser(&record); err != nil {
		log.Println("Err in parsing request.")
	}

	result := database.DBConn.Create(record)
	if result.Error != nil {
		log.Println("Err in saving data")
	}
	context["data"] = record

	c.Status(201)
	return c.JSON(context)

}

func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{}

	id := c.Params("id")
	var record model.Blog

	database.DBConn.First(&record, id)
	if record.ID == 0 {
		log.Println("Record not found")
		return c.JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Err in parsing request")
	}

	result := database.DBConn.Save(record)
	if result.Error != nil {
		log.Println("Err in saving data")
	}

	context["msg"] = "Update record successful"
	context["statusText"] = "OK"
	context["data"] = record

	c.Status(200)
	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "Delete a blog",
		"statusText": "OK",
	}

	id := c.Params("id")
	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not found")
		context["msg"] = "Record not found"

		c.Status(400)
		return c.JSON(context)
	}

	result := database.DBConn.Delete(record)

	if result.Error != nil {
		context["msg"] = "wrong delete in db"

		c.Status(400)
		return c.JSON(context)
	}

	context["msg"] = "Delete record successfull"
	c.Status(200)
	return c.JSON(context)
}
