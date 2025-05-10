package serializers

type OfficeDetailsInput struct {
	EmployeeCode string `json:"employee_code" binding:"required,max=10"`
	Address      string `json:"address" binding:"required,max=100"`
	City         string `json:"city" binding:"required,max=100"`
	State        string `json:"state" binding:"required,max=100"`
	Country      string `json:"country" binding:"required,max=100"`
	ContactNo    string `json:"contact_no" binding:"required,e164"`
	Email        string `json:"email" binding:"required,email,max=254"`
	Name         string `json:"name" binding:"required,max=100"`
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
