package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/marco210/blog-app/database"
	"github.com/marco210/blog-app/router"
)

func init()  {
	if err := godotenv.Load(".env"); err !=nil{
		log.Fatal("Err load .env file")
	}

	database.ConnectDB()
}

func main()  {
	sqlDb, err := database.DBConn.DB()

	if err!=nil{
		panic("Err in sql connection")
	}

	defer sqlDb.Close()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders:  "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New())

	router.SetupRoutes(app)
	
	app.Listen(":8080")
}