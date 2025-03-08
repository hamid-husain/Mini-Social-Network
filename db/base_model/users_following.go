package base_model

import "gorm.io/gorm"

type UserFollowing struct {
	gorm.Model
	UserID     uint `gorm:"primary key"`
	FollowerID uint `gorm:"primary key"`
}
