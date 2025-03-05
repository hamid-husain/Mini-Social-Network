package base_model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email              string              `gorm:"type:varchar(254);uniqueIndex:uniq_email,WHERE:deleted_at IS NULL; not null"`
	FirstName          string              `gorm:"type:varchar(46); not null"`
	LastName           string              `gorm:"type:varchar(46)"`
	DateOfBirth        string              `gorm:"type:date; not null"`
	Gender             uint8               `gorm:"type:smallint;check:gender IN (1,2,3)"`
	MaritalStatus      uint8               `gorm:"type:smallint;check:marital_status IN (1,2)"`
	Password           string              `gorm:"type:varchar(255); not null"`
	OfficeDetails      []OfficeDetail      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ResidentialDetails []ResidentialDetail `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Followers          []User              `gorm:"many2many:user_followers;joinForeignKey:UserID;joinReferences:FollowerID"`
}
