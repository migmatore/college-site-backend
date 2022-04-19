package service

import (
	"github.com/migmatore/college-site-backend/internal/core"
)

type SheduleStorage interface {
	Insert(shedule *core.Shedule)
	Get(shedule *core.Shedule)
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

func (s *SheduleService) Get(shedule *core.Shedule) {
	s.storage.Get(shedule)
}
