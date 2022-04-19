package storage

import (
	"github.com/migmatore/college-site-backend/internal/core"
	"gorm.io/gorm"
)

type SheduleStorage struct {
	db *gorm.DB
}

func NewSheduleStorage(db *gorm.DB) *SheduleStorage {
	return &SheduleStorage{db}
}

func (gs *SheduleStorage) Insert(shedule *core.Shedule) {
	gs.db.Create(&shedule)
}

func (gs *SheduleStorage) Get(shedule *core.Shedule) {
	gs.db.Preload("Group").Preload("Weekday").Preload("Subject").Preload("Office").Preload("Teacher").First(&shedule)
}
