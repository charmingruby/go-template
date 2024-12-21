package endpoint

import "github.com/gofiber/fiber/v2"

func (h *Handler) findAll(ctx *fiber.Ctx) error {
	users, err := h.service.FindAllUsers()

	if err != nil {
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(users)
}
