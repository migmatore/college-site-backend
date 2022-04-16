package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/college-site-backend/internal/core"
)

type GroupService interface {
	Create(group *core.Group) error
}

type GroupHandler struct {
	service GroupService
}

func NewGroupHandler(s GroupService) *GroupHandler {
	return &GroupHandler{service: s}
}

func (h *GroupHandler) Get(c *fiber.Ctx) error {
	return c.SendString("hello from group handler")
}

func (h *GroupHandler) Create(c *fiber.Ctx) error {
	h.service.Create(&core.Group{
		GroupName: "п-24",
	})

	return c.SendString("group was created")
}
