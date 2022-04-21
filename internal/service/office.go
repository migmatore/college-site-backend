package service

import (
	"strconv"

	"github.com/migmatore/college-site-backend/internal/core"
)

type OfficeStorage interface {
	Create(office *core.Office)
	GetAll(offices *[]core.Office)
	GetById(id int) *core.Office
	DeleteById(id int)
}

type OfficeService struct {
	storage OfficeStorage
}

func NewOfficeService(s OfficeStorage) *OfficeService {
	return &OfficeService{storage: s}
}

// Create new office
func (s *OfficeService) Create(officeNumber string) error {
	s.storage.Create(&core.Office{
		OfficeNumber: officeNumber,
	})

	return nil
}

// Get all offices
func (s *OfficeService) GetAll(offices *[]core.Office) {
	s.storage.GetAll(offices)
}

// Get office by id
func (s *OfficeService) GetById(id string) *core.Office {
	_id, _ := strconv.Atoi(id)

	return s.storage.GetById(_id)
}

// Delete office by id
func (s *OfficeService) DeleteById(id string) {
	_id, _ := strconv.Atoi(id)

	s.storage.DeleteById(_id)
}
