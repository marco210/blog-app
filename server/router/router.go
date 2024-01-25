package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marco210/blog-app/controller"
)

func SetupRoutes(app *fiber.App)  {
	app.Get("/", controller.BlogList)
	app.Get("/:id", controller.BlogListById)
	app.Post("/",controller.BlogCreate)
	app.Put("/:id", controller.BlogUpdate)
	app.Delete("/:id",controller.BlogDelete)
}