package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Deps struct {
	GroupService   GroupService
	OfficeService  OfficeService
	SubjectService SubjectService
	TeacherService TeacherService
	WeekdayService WeekdayService
	SheduleService SheduleService
}

type Handler struct {
	app     *fiber.App
	Group   *GroupHandler
	Office  *OfficeHandler
	Subject *SubjectHandler
	Teacher *TeacherHandler
	Weekday *WeekdayHandler
	Shedule *SheduleHandler
}

func New(deps Deps) *Handler {
	return &Handler{
		Group:   NewGroupHandler(deps.GroupService),
		Office:  NewOfficeHandler(deps.OfficeService),
		Subject: NewSubjectHandler(deps.SubjectService),
		Teacher: NewTeacherHandler(deps.TeacherService),
		Weekday: NewWeekdayHandler(deps.WeekdayService),
		Shedule: NewSheduleHandler(deps.SheduleService),
	}
}

func (h *Handler) Init() *fiber.App {
	h.app = fiber.New()

	h.app.Use(cors.New())

	api := h.app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/groups", h.Group.GetAll)           // was tested
	v1.Get("/group/:id", h.Group.GetById)       // was tested
	v1.Post("/group", h.Group.Create)           // was tested
	v1.Delete("/group/:id", h.Group.DeleteById) // was tested

	v1.Get("/offices", h.Office.GetAll)           // was tested
	v1.Get("/office/:id", h.Office.GetById)       // was tested
	v1.Post("/office", h.Office.Create)           // was tested
	v1.Delete("/office/:id", h.Office.DeleteById) // was tested

	v1.Get("/subjects", h.Subject.GetAll)           // was tested
	v1.Get("/subject/:id", h.Subject.GetById)       // was tested
	v1.Post("/subject", h.Subject.Create)           // was tested
	v1.Delete("/subject/:id", h.Subject.DeleteById) // was tested

	v1.Get("/teachers", h.Teacher.GetAll)           // was tested
	v1.Get("/teacher/:id", h.Teacher.GetById)       // was tested
	v1.Post("/teacher", h.Teacher.Create)           // was tested
	v1.Delete("/teacher/:id", h.Teacher.DeleteById) // was tested

	v1.Get("/weekdays", h.Weekday.GetAll)           // was tested
	v1.Get("/weekday/:id", h.Weekday.GetById)       // was tested
	v1.Post("/weekday", h.Weekday.Create)           // was tested
	v1.Delete("/weekday/:id", h.Weekday.DeleteById) // was tested

	v1.Get("/shedule", h.Shedule.Get) // was tested
	//v1.Get("/shedule/:id", h.Weekday.GetById)       // was tested
	v1.Post("/shedule", h.Shedule.Create) // was tested
	//v1.Delete("/shedule/:id", h.Weekday.DeleteById) // was tested

	return h.app
}
