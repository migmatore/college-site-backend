package storage

import (
	"github.com/migmatore/college-site-backend/internal/core"
	"gorm.io/gorm"
)

type OfficeStorage struct {
	db *gorm.DB
}

func NewOfficeStorage(db *gorm.DB) *OfficeStorage {
	return &OfficeStorage{db: db}
}

// Create new office
func (os *OfficeStorage) Create(office *core.Office) {
	os.db.Create(&office)
}

// Get all offices
func (os *OfficeStorage) GetAll(offices *[]core.Office) {
	os.db.Find(&offices)
}

// Get office by id
func (os *OfficeStorage) GetById(id int) *core.Office {
	var office core.Office

	os.db.First(&office, id)

	return &office
}

// Delete office by id
func (os *OfficeStorage) DeleteById(id int) {
	os.db.Delete(&core.Office{}, id)
}
