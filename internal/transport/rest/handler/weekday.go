package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/college-site-backend/internal/core"
	"github.com/migmatore/college-site-backend/pkg/util"
)

type WeekdayService interface {
	Create(weekdayName string) error
	GetAll(weekdays *[]core.Weekday)
	GetById(id string) *core.Weekday
	DeleteById(id string)
}

type WeekdayHandler struct {
	service WeekdayService
}

func NewWeekdayHandler(s WeekdayService) *WeekdayHandler {
	return &WeekdayHandler{service: s}
}

// Create new weekday
func (h *WeekdayHandler) Create(c *fiber.Ctx) error {
	p := new(core.Weekday)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	h.service.Create(p.WeekdayName)

	return util.NewResponse(c, fiber.StatusOK, "create weekday", "weekday was created")
}

// Get all weekdays
func (h *WeekdayHandler) GetAll(c *fiber.Ctx) error {
	var weekdays []core.Weekday

	h.service.GetAll(&weekdays)

	return util.NewResponse(c, fiber.StatusOK, "get all weekdays", weekdays)
}

// Get weekday by id
func (h *WeekdayHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	weekday := h.service.GetById(id)

	return util.NewResponse(c, fiber.StatusOK, "get weekday by id", weekday)
}

// Delete weekday by id
func (h *WeekdayHandler) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")

	h.service.DeleteById(id)

	return util.NewResponse(c, fiber.StatusOK, "delete weekday by id", fmt.Sprintf("weekday %s was deleted", id))
}
