package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marco210/blog-app/controller"
)

func SetupRoutes(app *fiber.App)  {
	app.Get("/", controller.BlogList)
	app.Post("/",controller.BlogCreate)
	app.Put("/", controller.BlogUpdate)
	app.Delete("/",controller.BlogDelete)
}