package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marco210/blog-app/database"
)

func init()  {
	database.ConnectDB()
}

func main()  {
	sqlDb, err := database.DBConn.DB()

	if err!=nil{
		panic("Err in sql connection")
	}

	defer sqlDb.Close()

	app := fiber.New()

	app.Get("/",func(c *fiber.Ctx) error {
		return c.JSON("done")
	})

	app.Get("/:name",func(c *fiber.Ctx) error {
		if c.Params("name") != ""{
			return c.SendString("Hello " +c.Params("name"))
		}

		return c.SendString("param missing name")
	})

	app.Listen(":8080")
}