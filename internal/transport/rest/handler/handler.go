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

	h.app.Get("/group", h.Group.Create)
	h.app.Get("/shedule", h.Shedule.Create)
	h.app.Get("/shedule-get", h.Shedule.Get)

	return h.app
}
