package core

type Shedule struct {
	Common
	Date string `json:"date" form:"date"`

	GroupId uint  `json:"group_id" form:"group_"`
	Group   Group `gorm:"foreignKey:GroupId"`

	WeekdayId uint    `json:"weekday_id" form:"weekday_id"`
	Weekday   Weekday `gorm:"foreignKey:WeekdayId"`

	SubjectId uint    `json:"subject_id" form:"subject_id"`
	Subject   Subject `gorm:"foreignKey:SubjectId"`

	OfficeId uint   `json:"office_id" form:"office_id"`
	Office   Office `gorm:"foreignKey:OfficeId"`

	TeacherId uint    `json:"teacher_id" form:"teacher_id"`
	Teacher   Teacher `gorm:"foreignKey:TeacherId"`
}
