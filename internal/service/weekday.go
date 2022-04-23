package service

import (
	"strconv"

	"github.com/migmatore/college-site-backend/internal/core"
)

type WeekdayStorage interface {
	Create(weekday *core.Weekday)
	GetAll(weekdays *[]core.Weekday)
	GetById(id int) *core.Weekday
	DeleteById(id int)
}

type WeekdayService struct {
	storage WeekdayStorage
}

func NewWeekdayService(s WeekdayStorage) *WeekdayService {
	return &WeekdayService{storage: s}
}

// Create new weekday
func (s *WeekdayService) Create(weekdayName string) error {
	s.storage.Create(&core.Weekday{
		WeekdayName: weekdayName,
	})

	return nil
}

// Get all weekday
func (s *WeekdayService) GetAll(weekdays *[]core.Weekday) {
	s.storage.GetAll(weekdays)
}

// Get weekday by id
func (s *WeekdayService) GetById(id string) *core.Weekday {
	_id, _ := strconv.Atoi(id)

	return s.storage.GetById(_id)
}

// Delete weekday by id
func (s *WeekdayService) DeleteById(id string) {
	_id, _ := strconv.Atoi(id)

	s.storage.DeleteById(_id)
}
