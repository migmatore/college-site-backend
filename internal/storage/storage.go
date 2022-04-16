package storage

import (
	"gorm.io/gorm"
)

type Storage struct {
	Group *GroupStorage
}

// Setup database connetion
func New(db *gorm.DB) *Storage {
	return &Storage{
		Group: NewGroupStorage(db),
	}
}

// Megrations setup and initialize
// func (storage *Storage) InitMigrations() {
// 	storage.DB.AutoMigrate(&core.Group{})
// }
