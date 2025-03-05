package base_model

type UserFollowing struct {
	UserID     uint `gorm:"primary key"`
	FollowerID uint `gorm:"primary key"`
}
