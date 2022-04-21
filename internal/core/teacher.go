package core

type Teacher struct {
	Common
	TeacherName string `json:"teacher_name" form:"teacher_name" gorm:"not null"`
}
