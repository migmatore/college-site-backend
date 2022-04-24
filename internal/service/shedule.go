package service

import (
	"strconv"

	"github.com/migmatore/college-site-backend/internal/core"
)

type SheduleStorage interface {
	Insert(shedule *core.Shedule)
	GetAll(shedule *[]core.Shedule)
	GetById(id int) *core.Shedule
	DeleteById(id int)
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

func (s *SheduleService) GetAll(shedule *[]core.Shedule) {
	s.storage.GetAll(shedule)
}

// Get shedule by id
func (s *SheduleService) GetById(id string) *core.Shedule {
	_id, _ := strconv.Atoi(id)

	return s.storage.GetById(_id)
}

// Delete shedule by id
func (s *SheduleService) DeleteById(id string) {
	_id, _ := strconv.Atoi(id)

	s.storage.DeleteById(_id)
}
