package http

import (
	"github.com/gofiber/fiber/v3"
	"net/http"
)

func OK(ctx fiber.Ctx, data interface{}) error {
	return ctx.Status(http.StatusOK).JSON(data)
}

func BadRequest(ctx fiber.Ctx, err error) error {
	return ctx.Status(http.StatusBadRequest).JSON(err)
}

func ServerError(ctx fiber.Ctx, err error) error {
	return ctx.Status(http.StatusInternalServerError).JSON(err)
}
