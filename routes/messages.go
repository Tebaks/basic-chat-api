package routes

import (
	"chatapp/controllers"
	"chatapp/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v3"
)

func MessagesRoute(route fiber.Router) {
	route.Use(jwt.New(jwt.Config{
		SigningKey: []byte("secret"),
	}))
	route.Use(middleware.JwtParseMiddleware)
	route.Get("/history", controllers.GetMessageHistory)
	route.Get("/historyByUser/:username", controllers.GetMessageHistoryByUser)
	route.Post("/send", controllers.SendMessage)
}
