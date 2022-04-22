package service

import (
	"strconv"

	"github.com/migmatore/college-site-backend/internal/core"
)

type TeacherStorage interface {
	Create(teacher *core.Teacher)
	GetAll(teachers *[]core.Teacher)
	GetById(id int) *core.Teacher
	DeleteById(id int)
}

type TeacherService struct {
	storage TeacherStorage
}

func NewTeacherService(s TeacherStorage) *TeacherService {
	return &TeacherService{storage: s}
}

// Create new teacher
func (s *TeacherService) Create(teacherName string) error {
	s.storage.Create(&core.Teacher{
		TeacherName: teacherName,
	})

	return nil
}

// Get all teachers
func (s *TeacherService) GetAll(teachers *[]core.Teacher) {
	s.storage.GetAll(teachers)
}

// Get teacher by id
func (s *TeacherService) GetById(id string) *core.Teacher {
	_id, _ := strconv.Atoi(id)

	return s.storage.GetById(_id)
}

// Delete teacher by id
func (s *TeacherService) DeleteById(id string) {
	_id, _ := strconv.Atoi(id)

	s.storage.DeleteById(_id)
}
