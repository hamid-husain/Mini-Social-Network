package serializers

type OfficeDetailsInput struct {
	EmployeeCode string `json:"employee_code" binding:"required"`
	Address      string `json:"address" binding:"required"`
	City         string `json:"city" binding:"required"`
	State        string `json:"state" binding:"required"`
	Country      string `json:"country" binding:"required"`
	ContactNo    string `json:"contact_no" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Name         string `json:"name" binding:"required"`
}

type OfficeDetailsResponse struct {
	ID           uint   `json:"id"`
	UserID       uint   `json:"user_id"`
	EmployeeCode string `json:"employee_code" binding:"required"`
	Address      string `json:"address" binding:"required"`
	City         string `json:"city" binding:"required"`
	State        string `json:"state" binding:"required"`
	Country      string `json:"country" binding:"required"`
	ContactNo    string `json:"contact_no" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Name         string `json:"name" binding:"required"`
}
