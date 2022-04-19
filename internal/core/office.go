package core

type Office struct {
	Common
	OfficeNumber uint32 `json:"office_number" gorm:"not null"`
}
