package routes

import (
	"chatapp/controllers"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v3"
)

func UsersRoute(route fiber.Router) {
	route.Post("/signup", controllers.SignUp)
	route.Post("/login", controllers.Login)
	route.Use(jwt.New(jwt.Config{
		SigningKey: []byte("secret"),
	}))
	route.Use(controllers.JwtParseMiddleware)
	route.Post("/block", controllers.BlockUser)
}
