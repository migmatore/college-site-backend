package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/college-site-backend/internal/core"
	"github.com/migmatore/college-site-backend/pkg/util"
)

type OfficeService interface {
	Create(officeNumber string) error
	GetAll(offices *[]core.Office)
	GetById(id string) *core.Office
	DeleteById(id string)
}

type OfficeHandler struct {
	service OfficeService
}

func NewOfficeHandler(s OfficeService) *OfficeHandler {
	return &OfficeHandler{service: s}
}

// Create new group
func (h *OfficeHandler) Create(c *fiber.Ctx) error {
	p := new(core.Office)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	h.service.Create(p.OfficeNumber)

	return util.NewResponse(c, fiber.StatusOK, "create office", "office was created")
}

// Get all groups
func (h *OfficeHandler) GetAll(c *fiber.Ctx) error {
	var offices []core.Office

	h.service.GetAll(&offices)

	return util.NewResponse(c, fiber.StatusOK, "get all offices", offices)
}

// Get group by id
func (h *OfficeHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	office := h.service.GetById(id)

	return util.NewResponse(c, fiber.StatusOK, "get office by id", office)
}

// Delete group by id
func (h *OfficeHandler) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")

	h.service.DeleteById(id)

	return util.NewResponse(c, fiber.StatusOK, "delete office by id", fmt.Sprintf("office %s was deleted", id))
}
