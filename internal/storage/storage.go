package storage

import (
	"github.com/migmatore/college-site-backend/internal/core"
	"gorm.io/gorm"
)

type Storage struct {
	Group   *GroupStorage
	Office  *OfficeStorage
	Subject *SubjectStorage
	Shedule *SheduleStorage
}

func New(db *gorm.DB) *Storage {
	db.AutoMigrate(&core.Shedule{})

	db.AutoMigrate(&core.Group{})
	db.AutoMigrate(&core.Office{})
	db.AutoMigrate(&core.Teacher{})
	db.AutoMigrate(&core.Weekday{})
	db.AutoMigrate(&core.Subject{})

	return &Storage{
		Group:   NewGroupStorage(db),
		Office:  NewOfficeStorage(db),
		Subject: NewSubjectStorage(db),
		Shedule: NewSheduleStorage(db),
	}
}
