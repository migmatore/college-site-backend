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

func (gs *GroupStorage) Insert(group *core.Group) {
	// gs.db.Create(&group)
}
