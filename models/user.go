package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email              string              `gorm:"type:varchar(254);uniqueIndex; not null"`
	FirstName          string              `gorm:"type:varchar(46); not null"`
	LastName           string              `gorm:"type:varchar(46)"`
	DateOfBirth        string              `gorm:"type:date; not null"`
	Gender             uint8               `gorm:"type:smallint;check:gender IN (1,2,3)"`
	MaritalStatus      uint8               `gorm:"type:smallint;check:marital_status IN (1,2)"`
	Password           string              `gorm:"type:varchar(255); not null"`
	OfficeDetails      []OfficeDetail      `gorm:"foreignKey:UserID"`
	ResidentialDetails []ResidentialDetail `gorm:"foreignKey:UserID"`
	Followers          []User              `gorm:"many2many:user_followers;joinForeignKey:UserID;joinReferences:FollowerID"`
}
