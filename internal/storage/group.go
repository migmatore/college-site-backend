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
	gs.db.Create(&group)
}

func (gs *GroupStorage) GetAll(groups *[]core.Group) {
	gs.db.Find(&groups)
}

func (gs *GroupStorage) GetById(id int) *core.Group {
	var g core.Group

	gs.db.First(&core.Group{}, id)

	return &g
}

func (gs *GroupStorage) DeleteById(id int) {
	gs.db.Delete(&core.Group{}, id)
}
