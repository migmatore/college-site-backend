package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/college-site-backend/internal/core"
)

type SheduleService interface {
	Create(shedule *core.Shedule) error
	Get(shedule *core.Shedule)
}

type SheduleHandler struct {
	service SheduleService
}

func NewSheduleHandler(s SheduleService) *SheduleHandler {
	return &SheduleHandler{service: s}
}

func (h *SheduleHandler) Create(c *fiber.Ctx) error {
	h.service.Create(&core.Shedule{
		Date: "10.0.10",
		Group: core.Group{
			GroupName: "p24",
		},
		Weekday: core.Weekday{
			WeekdayName: "monday",
		},
		Subject: core.Subject{
			SubjectName: "Networks",
		},
		Office: core.Office{
			OfficeNumber: "15",
		},
		Teacher: core.Teacher{
			TeacherName: "annannanna",
		},
	})

	return c.SendString("group was created")
}

func (h *SheduleHandler) Get(c *fiber.Ctx) error {
	var shedule core.Shedule

	h.service.Get(&shedule)

	return c.JSON(shedule)
}
