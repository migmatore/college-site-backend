package core

type Group struct {
	Common
	GroupName string `json:"group_name" form:"group_name" gorm:"not null"`
}
