package responses

import (
	"github.com/anhgeeky/go-temporal-labs/core/models"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Message string
}

func NotFound(c *fiber.Ctx) error {
	res := ErrorResponse{Message: "Endpoint not found"}
	return c.Status(fiber.StatusNotFound).JSON(res)
}

func WriteError(c *fiber.Ctx, err error) error {
	res := ErrorResponse{Message: err.Error()}
	return c.Status(fiber.StatusInternalServerError).JSON(res)
}

func SuccessResult[T any](c *fiber.Ctx, data T) error {
	return c.Status(fiber.StatusOK).JSON(models.Response[T]{
		Code: fiber.StatusOK,
		Data: data,
	})
}

func CreatedResult[T any](c *fiber.Ctx, data T) error {
	return c.Status(fiber.StatusCreated).JSON(models.Response[T]{
		Code: fiber.StatusOK,
		Data: data,
	})
}
