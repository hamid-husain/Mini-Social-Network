package users

import "time"

type OfficeDetail struct {
	ID            uint   `gorm:"primarykey"`
	UserID        uint   `json:"user_id" gorm:"not null;index"`
	EmployeeCode  string `json:"employee_code"`
	Address       string `json:"address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ContactNumber string `json:"contact_number"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
