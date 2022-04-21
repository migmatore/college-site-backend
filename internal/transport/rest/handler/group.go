package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/college-site-backend/internal/core"
	"github.com/migmatore/college-site-backend/pkg/util"
)

type GroupService interface {
	Create(groupName string) error
	GetAll(groups *[]core.Group)
	GetById(id string) *core.Group
	DeleteById(id string)
}

type GroupHandler struct {
	service GroupService
}

func NewGroupHandler(s GroupService) *GroupHandler {
	return &GroupHandler{service: s}
}

func (h *GroupHandler) Create(c *fiber.Ctx) error {
	p := new(core.Group)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	h.service.Create(p.GroupName)

	return util.NewResponse(c, fiber.StatusOK, "create group", "group was created")
}

func (h *GroupHandler) GetAll(c *fiber.Ctx) error {
	var groups []core.Group

	h.service.GetAll(&groups)

	return util.NewResponse(c, fiber.StatusOK, "get all groups", groups)
}

func (h *GroupHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	groups := h.service.GetById(id)

	return util.NewResponse(c, fiber.StatusOK, "get group by id", groups)
}

func (h *GroupHandler) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")

	h.service.DeleteById(id)

	return util.NewResponse(c, fiber.StatusOK, "delete group by id", fmt.Sprintf("group %s was deleted", id))
}
