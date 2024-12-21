package endpoint

import (
	"go-template/internal/identity/core/service"

	"github.com/gofiber/fiber/v2"
)

type findOneUserDTO struct {
	ID string `json:"id"`
}

func (h *Handler) findOne(ctx *fiber.Ctx) error {
	dto := findOneUserDTO{
		ID: ctx.Params("id"),
	}

	user, err := h.service.FindOneUser(service.FindOneUserInput{
		ID: dto.ID,
	})

	if err != nil {
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(user)
}
