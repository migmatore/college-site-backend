package storage

import (
	"github.com/migmatore/college-site-backend/internal/core"
	"gorm.io/gorm"
)

type GroupStorage struct {
	db *gorm.DB
}

func NewGroupStorage(db *gorm.DB) *GroupStorage {
	return &GroupStorage{db}
}

// Create new group
func (gs *GroupStorage) Create(group *core.Group) {
	gs.db.Create(&group)
}

// Get all groups
func (gs *GroupStorage) GetAll(groups *[]core.Group) {
	gs.db.Find(&groups)
}

// Get group by id
func (gs *GroupStorage) GetById(id int) *core.Group {
	var g core.Group

	gs.db.First(&core.Group{}, id)

	return &g
}

// Delete group by id
func (gs *GroupStorage) DeleteById(id int) {
	gs.db.Delete(&core.Group{}, id)
}
