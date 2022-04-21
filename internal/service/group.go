package service

import (
	"strconv"

	"github.com/migmatore/college-site-backend/internal/core"
)

type GroupStorage interface {
	Create(group *core.Group)
	GetAll(groups *[]core.Group)
	GetById(id int) *core.Group
	DeleteById(id int)
}

type GroupService struct {
	storage GroupStorage
}

func NewGroupService(storage GroupStorage) *GroupService {
	return &GroupService{storage: storage}
}

// Create new group
func (s *GroupService) Create(groupName string) error {
	s.storage.Create(&core.Group{
		GroupName: groupName,
	})

	return nil
}

// Get all groups
func (s *GroupService) GetAll(groups *[]core.Group) {
	s.storage.GetAll(groups)
}

// Get group by id
func (s *GroupService) GetById(id string) *core.Group {
	_id, _ := strconv.Atoi(id)

	return s.storage.GetById(_id)
}

// Delete goroup by id
func (s *GroupService) DeleteById(id string) {
	_id, _ := strconv.Atoi(id)

	s.storage.DeleteById(_id)
}
