package base_model

import "gorm.io/gorm"

type OfficeDetail struct {
	gorm.Model
	UserID        uint   `gorm:"index"`
	EmployeeCode  string `gorm:"type:varchar(10)"`
	Address       string `gorm:"type:varchar(255)"`
	City          string `gorm:"type:varchar(50)"`
	State         string `gorm:"type:varchar(50)"`
	Country       string `gorm:"type:varchar(100)"`
	ContactNumber string `gorm:"type:varchar(15)"`
	Email         string `gorm:"type:varchar(254)"`
	Name          string `gorm:"type:varchar(100)"`
}
