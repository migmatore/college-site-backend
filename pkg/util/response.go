package util

import "github.com/gofiber/fiber/v2"

func NewResponse(c *fiber.Ctx, statusCode int, request string, response interface{}) error {
	return c.JSON(fiber.Map{
		"status":   statusCode,
		"request":  request,
		"response": response,
	})
}
