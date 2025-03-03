package models

type UserFollowing struct {
	UserID     uint `gorm:"primary key"`
	FollowerID uint `gorm:"primary key"`
}
