package service

import (
	"github.com/migmatore/college-site-backend/internal/core"
)

type SheduleStorage interface {
	Insert(shedule *core.Shedule)
	Get(shedule *[]core.Shedule)
	Find(id uint) *core.Group
}

type SheduleService struct {
	storage SheduleStorage
}

func NewSheduleService(storage SheduleStorage) *SheduleService {
	return &SheduleService{storage: storage}
}

func (s *SheduleService) Create(shedule *core.Shedule) error {
	s.storage.Insert(shedule)

	return nil
}

func (s *SheduleService) Get(shedule *[]core.Shedule) {
	s.storage.Get(shedule)
}

func (s *SheduleService) Find(id uint) *core.Group {
	return s.storage.Find(id)
}
