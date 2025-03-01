package baseModel

import (
	"gorm.io/gorm"
)

type ResidentialDetail struct {
	gorm.Model
	UserID         uint   `gorm:"index"`
	Address        string `gorm:"type:varchar(100)"`
	City           string `gorm:"type:varchar(100)"`
	State          string `gorm:"type:varchar(100)"`
	Country        string `gorm:"type:varchar(100)"`
	ContactNumber1 string `gorm:"type:varchar(15)"`
	ContactNumber2 string `gorm:"type:varchar(15)"`
}
