package service

import "github.com/migmatore/college-site-backend/internal/core"

type GroupStorage interface {
	Insert(group *core.Group)
}

type GroupService struct {
	storage GroupStorage
}

func NewGroupService(storage GroupStorage) *GroupService {
	return &GroupService{storage: storage}
}

func (s *GroupService) Create(group *core.Group) error {
	s.storage.Insert(group)

	return nil
}
