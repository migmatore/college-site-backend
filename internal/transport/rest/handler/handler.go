package handler

import "github.com/gofiber/fiber/v2"

type Deps struct {
	GroupService   GroupService
	SheduleService SheduleService
}

type Handler struct {
	app     *fiber.App
	Group   *GroupHandler
	Shedule *SheduleHandler
}

func New(deps Deps) *Handler {
	return &Handler{
		Group:   NewGroupHandler(deps.GroupService),
		Shedule: NewSheduleHandler(deps.SheduleService),
	}
}

func (h *Handler) Init() *fiber.App {
	h.app = fiber.New()

	api := h.app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/groups", h.Group.GetAll)           // was tested
	v1.Get("/group/:id", h.Group.GetById)       // was tested
	v1.Get("/group-create", h.Group.Create)     // was tested
	v1.Delete("/group/:id", h.Group.DeleteById) // was tested
	v1.Post("/group", h.Group.Create)           // was tested

	return h.app
}
