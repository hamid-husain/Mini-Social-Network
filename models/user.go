package models

import (
	"strings"

	"gorm.io/gorm"

	"fmt"
	"mini-social-network/constants"
)

type User struct {
	gorm.Model
	Email              string              `gorm:"type:varchar(254);uniqueIndex:uniq_email,WHERE:deleted_at IS NULL; not null"`
	FirstName          string              `gorm:"type:varchar(60); not null"`
	LastName           string              `gorm:"type:varchar(60)"`
	DateOfBirth        string              `gorm:"type:date; not null"`
	Gender             uint8               `gorm:"type:smallint; not null"`
	MaritalStatus      uint8               `gorm:"type:smallint; not null"`
	Password           string              `gorm:"type:varchar(255); not null"`
	OfficeDetails      []OfficeDetail      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ResidentialDetails []ResidentialDetail `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Following          []UserFollowing     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Followers          []UserFollowing     `gorm:"foreignKey:FollowerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (u *User) SetGender(genderStr string) error {
	genderStr = strings.ToLower(genderStr)
	switch genderStr {
	case constants.Male:
		u.Gender = constants.GenderMale
	case constants.Female:
		u.Gender = constants.GenderFemale
	case constants.Other:
		u.Gender = constants.GenderOther
	default:
		return fmt.Errorf("invalid gender value: %s", genderStr)
	}
	return nil
}

func (u *User) SetMaritalStatus(maritalStatus string) error {
	maritalStatus = strings.ToLower(maritalStatus)
	switch maritalStatus {
	case constants.Single:
		u.MaritalStatus = constants.MaritalStatusSingle
	case constants.Married:
		u.MaritalStatus = constants.MaritalStatusMarried
	default:
		return fmt.Errorf("invalid marital_status value: %s", maritalStatus)
	}

	return nil
}
