package service

import (
	"strconv"

	"github.com/migmatore/college-site-backend/internal/core"
)

type GroupStorage interface {
	Insert(group *core.Group)
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

func (s *GroupService) Create(groupName string) error {
	s.storage.Insert(&core.Group{
		GroupName: groupName,
	})

	return nil
}

func (s *GroupService) GetAll(groups *[]core.Group) {
	s.storage.GetAll(groups)
}

func (s *GroupService) GetById(id string) *core.Group {
	_id, _ := strconv.Atoi(id)

	return s.storage.GetById(_id)
}

func (s *GroupService) DeleteById(id string) {
	_id, _ := strconv.Atoi(id)

	s.storage.DeleteById(_id)
}
