package service

import (
	"strconv"

	"github.com/migmatore/college-site-backend/internal/core"
)

type SubjectStorage interface {
	Create(subject *core.Subject)
	GetAll(subjects *[]core.Subject)
	GetById(id int) *core.Subject
	DeleteById(id int)
}

type SubjectService struct {
	storage SubjectStorage
}

func NewSubjectService(s SubjectStorage) *SubjectService {
	return &SubjectService{storage: s}
}

// Create new subject
func (s *SubjectService) Create(subjectName string) error {
	s.storage.Create(&core.Subject{
		SubjectName: subjectName,
	})

	return nil
}

// Get all subjects
func (s *SubjectService) GetAll(subjects *[]core.Subject) {
	s.storage.GetAll(subjects)
}

// Get subject by id
func (s *SubjectService) GetById(id string) *core.Subject {
	_id, _ := strconv.Atoi(id)

	return s.storage.GetById(_id)
}

// Delete subject by id
func (s *SubjectService) DeleteById(id string) {
	_id, _ := strconv.Atoi(id)

	s.storage.DeleteById(_id)
}
