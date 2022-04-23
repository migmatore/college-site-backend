package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/college-site-backend/internal/core"
)

type SheduleService interface {
	Create(shedule *core.Shedule) error
	Get(shedule *[]core.Shedule)
	Find(id uint) *core.Group
}

type SheduleHandler struct {
	service SheduleService
}

func NewSheduleHandler(s SheduleService) *SheduleHandler {
	return &SheduleHandler{service: s}
}

func (h *SheduleHandler) Create(c *fiber.Ctx) error {
	s := new(core.Shedule)

	if err := c.BodyParser(s); err != nil {
		return err
	}

	//group := h.service.Find(p.GroupId)

	h.service.Create(&core.Shedule{
		Date:      s.Date,
		GroupId:   s.GroupId,
		TeacherId: s.TeacherId,
		WeekdayId: s.WeekdayId,
		SubjectId: s.SubjectId,
		OfficeId:  s.OfficeId,
	})

	return c.SendString("shedule was created")
}

func (h *SheduleHandler) Get(c *fiber.Ctx) error {
	var shedule []core.Shedule

	h.service.Get(&shedule)

	return c.JSON(shedule)
}
