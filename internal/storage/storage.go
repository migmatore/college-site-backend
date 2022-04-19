package storage

import (
	"github.com/migmatore/college-site-backend/internal/core"
	"gorm.io/gorm"
)

type Storage struct {
	Group   *GroupStorage
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
		Shedule: NewSheduleStorage(db),
	}
}

// Megrations setup and initialize
// func (storage *Storage) InitMigrations() {
// 	storage.DB.AutoMigrate(&core.Group{})
// }
