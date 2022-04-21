package handler

import "github.com/gofiber/fiber/v2"

type Deps struct {
	GroupService   GroupService
	OfficeService  OfficeService
	SheduleService SheduleService
}

type Handler struct {
	app     *fiber.App
	Group   *GroupHandler
	Office  *OfficeHandler
	Shedule *SheduleHandler
}

func New(deps Deps) *Handler {
	return &Handler{
		Group:   NewGroupHandler(deps.GroupService),
		Office:  NewOfficeHandler(deps.OfficeService),
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

	v1.Get("/offices", h.Office.GetAll)           // was tested
	v1.Get("/office/:id", h.Office.GetById)       // was tested
	v1.Get("/office-create", h.Office.Create)     // was tested
	v1.Delete("/office/:id", h.Office.DeleteById) // was tested
	v1.Post("/office", h.Office.Create)           // was tested

	return h.app
}
