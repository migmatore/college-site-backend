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

func (ss *SheduleStorage) GetAll(shedule *[]core.Shedule) {
	ss.db.Preload("Group").Preload("Weekday").Preload("Subject").Preload("Office").Preload("Teacher").Find(&shedule)
}

// Get shedule by id
func (os *SheduleStorage) GetById(id int) *core.Shedule {
	var shedule core.Shedule

	os.db.First(&shedule, id)

	return &shedule
}

// Delete shedule by id
func (os *SheduleStorage) DeleteById(id int) {
	os.db.Delete(&core.Shedule{}, id)
}
