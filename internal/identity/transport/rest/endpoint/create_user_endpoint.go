package endpoint

import (
	"errors"
	"go-template/internal/identity/core/service"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

type createUserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *Handler) create(ctx *fiber.Ctx) error {
	dto := new(createUserDTO)

	if err := ctx.BodyParser(dto); err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "bad request"})
	}

	// TODO: move this validation to a separate structure
	if err := func() error {
		if dto.Name == "" {
			return errors.New("missing name")
		}

		if len(dto.Name) > 100 {
			return errors.New("invalid name")
		}

		if dto.Email == "" {
			return errors.New("missing email")
		}

		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		re := regexp.MustCompile(emailRegex)
		if !re.MatchString(dto.Email) {
			return errors.New("invalid email")
		}
		return nil
	}(); err != nil {
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.CreateUser(service.CreateUserInput{
		Name:  dto.Name,
		Email: dto.Email,
	}); err != nil {
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.
		Status(fiber.StatusOK).
		Send([]byte{})
}
