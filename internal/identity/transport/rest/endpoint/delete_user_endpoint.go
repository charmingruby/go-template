package endpoint

import (
	"go-template/internal/identity/core/service"

	"github.com/gofiber/fiber/v2"
)

type deleteUserDTO struct {
	ID string `json:"id"`
}

func (h *Handler) delete(ctx *fiber.Ctx) error {
	dto := new(deleteUserDTO)

	if err := ctx.BodyParser(dto); err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "bad request"})
	}

	if err := h.service.DeleteUser(service.DeleteUserInput{
		ID: dto.ID,
	}); err != nil {
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.
		Status(fiber.StatusOK).
		Send([]byte{})
}
