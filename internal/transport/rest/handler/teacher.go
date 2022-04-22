package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/college-site-backend/internal/core"
	"github.com/migmatore/college-site-backend/pkg/util"
)

type TeacherService interface {
	Create(teacherName string) error
	GetAll(teachers *[]core.Teacher)
	GetById(id string) *core.Teacher
	DeleteById(id string)
}

type TeacherHundler struct {
	service TeacherService
}

func NewTeacherHundler(s TeacherService) *TeacherHundler {
	return &TeacherHundler{service: s}
}

// Create new teacher
func (h *TeacherHundler) Create(c *fiber.Ctx) error {
	p := new(core.Teacher)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	h.service.Create(p.TeacherName)

	return util.NewResponse(c, fiber.StatusOK, "create teacher", "teacher was created")
}

// Get all teachers
func (h *TeacherHundler) GetAll(c *fiber.Ctx) error {
	var teachers []core.Teacher

	h.service.GetAll(&teachers)

	return util.NewResponse(c, fiber.StatusOK, "get all teachers", teachers)
}

// Get teacher by id
func (h *TeacherHundler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	teacher := h.service.GetById(id)

	return util.NewResponse(c, fiber.StatusOK, "get teacher by id", teacher)
}

// Delete teacher by id
func (h *TeacherHundler) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")

	h.service.DeleteById(id)

	return util.NewResponse(c, fiber.StatusOK, "delete teacher by id", fmt.Sprintf("teacher %s was deleted", id))
}
