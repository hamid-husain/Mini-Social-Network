package base_model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email              string              `gorm:"type:varchar(254);uniqueIndex:uniq_email,WHERE:deleted_at IS NULL; not null"`
	FirstName          string              `gorm:"type:varchar(46); not null"`
	LastName           string              `gorm:"type:varchar(46)"`
	DateOfBirth        string              `gorm:"type:date; not null"`
	Gender             uint8               `gorm:"type:smallint; not null"`
	MaritalStatus      uint8               `gorm:"type:smallint; not null"`
	Password           string              `gorm:"type:varchar(255); not null"`
	OfficeDetails      []OfficeDetail      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ResidentialDetails []ResidentialDetail `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Following          []UserFollowing     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Followers          []UserFollowing     `gorm:"foreignKey:FollowerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
