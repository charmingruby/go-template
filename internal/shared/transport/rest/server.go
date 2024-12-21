package rest

import (
	"github.com/gofiber/fiber/v2"
)

func New() (app *fiber.App) {
	app = fiber.New()

	app.Use("*", func(ctx *fiber.Ctx) error {
		ctx.Set("Content-Type", "application/json")

		return ctx.Next()
	})

	return app
}
