package storage

import (
	"github.com/migmatore/college-site-backend/internal/core"
	"gorm.io/gorm"
)

type Storage struct {
	Group *GroupStorage
}

func New(db *gorm.DB) *Storage {
	db.AutoMigrate(&core.Group{})

	return &Storage{
		Group: NewGroupStorage(db),
	}
}

// Megrations setup and initialize
// func (storage *Storage) InitMigrations() {
// 	storage.DB.AutoMigrate(&core.Group{})
// }
