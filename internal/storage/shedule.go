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

func (ss *SheduleStorage) Insert(shedule *core.Shedule) {
	ss.db.Create(&shedule)
}

func (ss *SheduleStorage) Get(shedule *[]core.Shedule) {
	ss.db.Preload("Group").Preload("Weekday").Preload("Subject").Preload("Office").Preload("Teacher").Find(&shedule)
}

func (ss *SheduleStorage) Find(id uint) *core.Group {
	var group core.Group

	ss.db.Find(&group, id)

	return &group
}
