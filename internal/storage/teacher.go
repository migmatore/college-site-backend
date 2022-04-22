package storage

import (
	"github.com/migmatore/college-site-backend/internal/core"
	"gorm.io/gorm"
)

type TeacherStorage struct {
	db *gorm.DB
}

func NewTeacherStorage(db *gorm.DB) *TeacherStorage {
	return &TeacherStorage{db: db}
}

// Create new teacher
func (gs *TeacherStorage) Create(teacher *core.Teacher) {
	gs.db.Create(&teacher)
}

// Get all teachers
func (gs *TeacherStorage) GetAll(teachers *[]core.Teacher) {
	gs.db.Find(&teachers)
}

// Get teacher by id
func (gs *TeacherStorage) GetById(id int) *core.Teacher {
	var teacher core.Teacher

	gs.db.First(&teacher, id)

	return &teacher
}

// Delete teacher by id
func (gs *TeacherStorage) DeleteById(id int) {
	gs.db.Delete(&core.Teacher{}, id)
}
