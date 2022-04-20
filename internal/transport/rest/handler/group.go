package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/college-site-backend/internal/core"
	"github.com/migmatore/college-site-backend/pkg/util"
)

type GroupService interface {
	Create(group *core.Group) error
	GetAll(groups []*core.Group)
	GetById(id string) *core.Group
}

type GroupHandler struct {
	service GroupService
}

func NewGroupHandler(s GroupService) *GroupHandler {
	return &GroupHandler{service: s}
}

func (h *GroupHandler) Create(c *fiber.Ctx) error {
	h.service.Create(&core.Group{
		GroupName: "п-24",
	})

	return c.SendString("group was created")
}

func (h *GroupHandler) GetAll(c *fiber.Ctx) error {
	var groups []core.Group

	h.service.GetAll(groups)

	return util.NewResponse(c, http.StatusOK, )

func (h *GroupHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	groups := h.service.GetById(id)

	return util.NewResponse(c, http.StatusOK, "get group by id", groups)
}
