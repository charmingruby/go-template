package endpoint

import (
	"go-template/internal/identity/core/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service service.Service
}

func New(svc service.Service) *Handler {
	return &Handler{
		service: svc,
	}
}

func (h *Handler) Setup(router *fiber.App) {
	router.Post("/user", h.create)
	router.Get("/user", h.findAll)
	router.Get("/user/:id", h.findOne)
	router.Put("/user", h.update)
	router.Delete("/user/:id", h.delete)
}
