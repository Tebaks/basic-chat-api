package controllers

import (
	"chatapp/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GetMessageHistory handles get user's messages request
func GetMessageHistory(c *fiber.Ctx) error {
	username := fmt.Sprintf("%v", c.Locals("username"))

	messages, err := models.GetUsersMessages(username)
	if err != nil {
		return ResponseCreator(c, 400, nil, false, "Failed to get messages", err.Error())
	}

	return ResponseCreator(c, 200, fiber.Map{"messages": messages}, true, "Get messages successfully", "")
}

// GetMEssageHistoryByUser handles get user's messages by spesific user request
func GetMessageHistoryByUser(c *fiber.Ctx) error {
	byUser := c.Params("username")
	username := fmt.Sprintf("%v", c.Locals("username"))

	messages, err := models.GetUsersMessagesByUser(username, byUser)
	if err != nil {
		return ResponseCreator(c, 400, nil, false, "Failed to get messages", err.Error())
	}

	return ResponseCreator(c, 200, fiber.Map{"messages": messages}, true, "Get messages successfully", "")
}

// SendMessage handles send message request
func SendMessage(c *fiber.Ctx) error {
	messageRequest := new(models.SendMessageRequest)

	if err := c.BodyParser(messageRequest); err != nil {
		return ResponseCreator(c, 400, nil, false, "Failed to parse body", err.Error())
	}

	username := fmt.Sprintf("%v", c.Locals("username"))

	err := models.SendMessage(username, messageRequest.To, messageRequest.Content)

	if err != nil {
		return ResponseCreator(c, 500, nil, false, "Error sending message", err.Error())
	}

	return ResponseCreator(c, 200, nil, true, "Message sended successfully", "")
}
