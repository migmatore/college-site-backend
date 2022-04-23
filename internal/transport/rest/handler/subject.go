package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/college-site-backend/internal/core"
	"github.com/migmatore/college-site-backend/pkg/util"
)

type SubjectService interface {
	Create(subjectName string) error
	GetAll(subjects *[]core.Subject)
	GetById(id string) *core.Subject
	DeleteById(id string)
}

type SubjectHandler struct {
	service SubjectService
}

func NewSubjectHandler(s SubjectService) *SubjectHandler {
	return &SubjectHandler{service: s}
}

// Create new subject
func (h *SubjectHandler) Create(c *fiber.Ctx) error {
	p := new(core.Subject)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	h.service.Create(p.SubjectName)

	return util.NewResponse(c, fiber.StatusOK, "create subject", "subject was created")
}

// Get all subjects
func (h *SubjectHandler) GetAll(c *fiber.Ctx) error {
	var subjects []core.Subject

	h.service.GetAll(&subjects)

	return util.NewResponse(c, fiber.StatusOK, "get all subjects", subjects)
}

// Get subject by id
func (h *SubjectHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	subjects := h.service.GetById(id)

	return util.NewResponse(c, fiber.StatusOK, "get subject by id", subjects)
}

// Delete subject by id
func (h *SubjectHandler) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")

	h.service.DeleteById(id)

	return util.NewResponse(c, fiber.StatusOK, "delete subject by id", fmt.Sprintf("subject %s was deleted", id))
}
