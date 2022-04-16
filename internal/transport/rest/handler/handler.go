package handler

import "github.com/gofiber/fiber/v2"

type Deps struct {
	GroupService GroupService
}

type Handler struct {
	app   *fiber.App
	Group *GroupHandler
}

func New(deps Deps) *Handler {
	return &Handler{
		Group: NewGroupHandler(deps.GroupService),
	}
}

func (h *Handler) Init() *fiber.App {
	h.app = fiber.New()

	h.app.Get("/group", h.Group.Create)

	return h.app
}
