package models

import "gorm.io/gorm"

type UserFollowing struct {
	UserID     uint           `gorm:"primary key"`
	FollowerID uint           `gorm:"primary key"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
