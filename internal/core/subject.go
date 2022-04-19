package core

type Subject struct {
	Common
	SubjectName string `json:"subject_name" gorm:"not null"`
}
