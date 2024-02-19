package apis

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	Message string
}

func NotFoundHandler(c *fiber.Ctx) error {
	res := ErrorResponse{Message: "Endpoint not found"}
	return c.Status(fiber.StatusNotFound).JSON(res)
}

func WriteError(c *fiber.Ctx, err error) error {
	res := ErrorResponse{Message: err.Error()}
	return c.Status(fiber.StatusInternalServerError).JSON(res)
}
