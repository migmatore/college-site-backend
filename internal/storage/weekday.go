package storage

import (
	"github.com/migmatore/college-site-backend/internal/core"
	"gorm.io/gorm"
)

type WeekdayStorage struct {
	db *gorm.DB
}

func NewWeekdayStorage(db *gorm.DB) *WeekdayStorage {
	return &WeekdayStorage{db: db}
}

// Create new weekday
func (gs *WeekdayStorage) Create(weekday *core.Weekday) {
	gs.db.Create(&weekday)
}

// Get all weekdays
func (gs *WeekdayStorage) GetAll(weekday *[]core.Weekday) {
	gs.db.Find(&weekday)
}

// Get weekday by id
func (gs *WeekdayStorage) GetById(id int) *core.Weekday {
	var weekday core.Weekday

	gs.db.First(&weekday, id)

	return &weekday
}

// Delete weekday by id
func (gs *WeekdayStorage) DeleteById(id int) {
	gs.db.Delete(&core.Weekday{}, id)
}
