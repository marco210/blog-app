package main

import (
	"github.com/gofiber/fiber/v2"
)

func main()  {
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