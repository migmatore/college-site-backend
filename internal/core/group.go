package core

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Id        uint   `json:"id" gorm:"not null; unique; autoIncrement"`
	GroupName string `json:"group_name gorm:"not null"`
}
