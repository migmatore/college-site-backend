package core

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Id        uint `gorm:"not null; unique; autoIncrement"`
	GroupName string
}
