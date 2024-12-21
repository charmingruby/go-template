package endpoint

import (
	"go-template/internal/identity/core/service"

	"github.com/gofiber/fiber/v2"
)

type updateUserDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *Handler) update(ctx *fiber.Ctx) error {
	dto := new(updateUserDTO)

	if err := ctx.BodyParser(dto); err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "bad request"})
	}

	user, err := h.service.UpdateUser(service.UpdateUserInput{
		ID:    dto.ID,
		Name:  dto.Name,
		Email: dto.Email,
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
