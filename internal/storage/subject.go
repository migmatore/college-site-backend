package storage

import (
	"github.com/migmatore/college-site-backend/internal/core"
	"gorm.io/gorm"
)

type SubjectStorage struct {
	db *gorm.DB
}

func NewSubjectStorage(db *gorm.DB) *SubjectStorage {
	return &SubjectStorage{db: db}
}

// Create new subject
func (gs *SubjectStorage) Create(subject *core.Subject) {
	gs.db.Create(&subject)
}

// Get all subjects
func (gs *SubjectStorage) GetAll(subjects *[]core.Subject) {
	gs.db.Find(&subjects)
}

// Get subject by id
func (gs *SubjectStorage) GetById(id int) *core.Subject {
	var subject core.Subject

	gs.db.First(&subject, id)

	return &subject
}

// Delete subject by id
func (gs *SubjectStorage) DeleteById(id int) {
	gs.db.Delete(&core.Subject{}, id)
}
