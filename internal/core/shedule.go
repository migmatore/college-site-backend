package core

type Shedule struct {
	Common
	Date      string `json:"date"`
	GroupId   uint
	Group     Group `json:"group_id" gorm:"foreignKey:GroupId"`
	WeekdayId uint
	Weekday   Weekday `json:"weekday_id" gorm:"foreignKey:WeekdayId"`
	SubjectId uint
	Subject   Subject `json:"subject_id" gorm:"foreignKey:SubjectId"`
	OfficeId  uint
	Office    Office `json:"office_id" gorm:"foreignKey:OfficeId"`
	TeacherId uint
	Teacher   Teacher `json:"teacher_id" gorm:"foreignKey:TeacherId"`
}
