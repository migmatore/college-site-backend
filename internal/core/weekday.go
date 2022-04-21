package core

type Weekday struct {
	Common
	WeekdayName string `json:"weekday_name" form:"weekday_name" gorm:"not null"`
}
