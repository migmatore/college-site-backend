package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/college-site-backend/internal/core"
	"github.com/migmatore/college-site-backend/pkg/util"
)

type SheduleService interface {
	Create(shedule *core.Shedule) error
	GetAll(shedule *[]core.Shedule)
	GetById(id string) *core.Shedule
	DeleteById(id string)
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

	return util.NewResponse(c, fiber.StatusOK, "create shedule", "shedule was created")
}

func (h *SheduleHandler) GetAll(c *fiber.Ctx) error {
	var shedule []core.Shedule

	h.service.GetAll(&shedule)

	return c.JSON(shedule)
}

func (h *SheduleHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	shedule := h.service.GetById(id)

	return util.NewResponse(c, fiber.StatusOK, "get shedule by id", shedule)
}

// Delete shedule by id
func (h *SheduleHandler) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")

	h.service.DeleteById(id)

	return util.NewResponse(c, fiber.StatusOK, "delete shedule by id", fmt.Sprintf("shedule %s was deleted", id))
}
