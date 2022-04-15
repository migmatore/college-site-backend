package storage

import (
	"github.com/migmatore/college-site-backend/internal/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

// Setup database connetion
func (storage *Storage) New(dsn string) (*Storage, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	_storage := Storage{
		db: db,
	}

	return &_storage, err
}

// Megrations setup and initialize
func (storage *Storage) InitMigrations() {
	storage.db.AutoMigrate(&core.Group{})
}
