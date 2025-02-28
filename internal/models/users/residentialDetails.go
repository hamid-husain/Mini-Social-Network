package users

import "time"

type ResidentialDetail struct {
	ID             uint   `gorm:"primarykey"`
	UserID         uint   `json:"user_id" gorm:"not null;index"`
	Address        string `json:"address"`
	City           string `json:"city"`
	State          string `json:"state"`
	Country        string `json:"country"`
	ContactNumber1 string `json:"contact_number_1"`
	ContactNumber2 string `json:"contact_number_2,omitempty"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
