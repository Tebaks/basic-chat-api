package controllers

import (
	"chatapp/models"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// BlockUser handles user's block request
func BlockUser(c *fiber.Ctx) error {
	blockRequest := new(models.BlockRequest)

	if err := c.BodyParser(blockRequest); err != nil {
		log.Println(err)
		return ResponseCreator(c, 400, nil, false, "Failed to parse body", err.Error())
	}
	username := fmt.Sprintf("%v", c.Locals("username"))
	err := models.BlockUser(username, blockRequest.Username)
	if err != nil {
		return ResponseCreator(c, 500, nil, false, "Failed to block user", err.Error())
	}

	return ResponseCreator(c, 200, nil, true, "Block user successfully", "")
}

// Login handles user's login request
func Login(c *fiber.Ctx) error {
	credentials := new(models.LoginRequest)

	if err := c.BodyParser(credentials); err != nil {
		log.Println(err)
		return ResponseCreator(c, 400, nil, false, "Failed to parse body", err.Error())
	}
	token, err := models.Login(credentials)
	if err != nil {
		return ResponseCreator(c, 500, nil, false, "Failed to login", err.Error())
	}

	return ResponseCreator(c, 200, fiber.Map{"token": token}, true, "User login successfully", "")
}

// SignUp handles user's signup request
func SignUp(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		log.Println(err)
		return ResponseCreator(c, 400, nil, false, "Failed to parse body", err.Error())
	}

	token, err := models.SignUp(user)
	if err != nil {
		return ResponseCreator(c, 500, nil, false, "Failed to signup", err.Error())
	}
	return ResponseCreator(c, 201, fiber.Map{"token": token}, true, "User signup successfully", "")
}

func ResponseCreator(c *fiber.Ctx, code int, data interface{}, success bool, message string, err string) error {
	return c.Status(code).JSON(fiber.Map{
		"data":    data,
		"success": success,
		"message": message,
		"error":   err,
	})
}
