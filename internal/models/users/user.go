package users

import (
	"time"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
	Other  Gender = "other"
)

type MaritalStatus string

const (
	Single  MaritalStatus = "single"
	Married MaritalStatus = "married"
)

type User struct {
	ID                uint          `gorm:"primarykey"`
	Email             string        `json:"email" gorm:"unique; not null"`
	FirstName         string        `json:"first_name"`
	LastName          string        `json:"last_name"`
	DateOfBirth       time.Time     `json:"date_of_birth"`
	Gender            Gender        `json:"gender"`
	MaritalStatus     MaritalStatus `json:"marital_status"`
	Password          string        `json:"password" gorm:"not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	OfficeDetail      []OfficeDetail      `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ResidentialDetail []ResidentialDetail `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
