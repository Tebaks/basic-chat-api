package main

import (
	"chatapp/config"
	"chatapp/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	routes.UsersRoute(api.Group("/user"))
	routes.MessagesRoute(api.Group("/message"))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	config.ConnectDB()
	setupRoutes(app)

	port := os.Getenv("PORT")
	err = app.Listen(":" + port)
	if err != nil {
		log.Fatal("Error app failed to start")
		panic(err)
	}
}
