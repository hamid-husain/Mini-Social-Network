package baseModel

type OfficeDetail struct {
	ID            uint   `gorm:"primaryKey"`
	UserID        uint   `gorm:"index"`
	EmployeeCode  string `gorm:"type:varchar(6)"`
	Address       string `gorm:"type:varchar(100)"`
	City          string `gorm:"type:varchar(100)"`
	State         string `gorm:"type:varchar(100)"`
	Country       string `gorm:"type:varchar(100)"`
	ContactNumber string `gorm:"type:varchar(15)"`
	Email         string `gorm:"type:varchar(254)"`
	Name          string `gorm:"type:varchar(100)"`
}
