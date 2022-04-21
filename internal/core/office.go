package core

type Office struct {
	Common
	OfficeNumber string `json:"office_number" form:"office_number" gorm:"not null"`
}
