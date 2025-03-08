package serializers

type ResidentialDetailsInput struct {
	Address    string `json:"address" binding:"required,max=255"`
	City       string `json:"city" binding:"required,max=50"`
	State      string `json:"state" binding:"required,max=50"`
	Country    string `json:"country" binding:"required,max=100"`
	ContactNo1 string `json:"residential_contact_no_1" binding:"required,e164"`
	ContactNo2 string `json:"residential_contact_no_2" binding:"omitempty,e164"`
}

type ResidentialDetailsResponse struct {
	ID         uint   `json:"id"`
	UserID     uint   `json:"user_id"`
	Address    string `json:"address" binding:"required"`
	City       string `json:"city" binding:"required"`
	State      string `json:"state" binding:"required"`
	Country    string `json:"country" binding:"required"`
	ContactNo1 string `json:"residential_contact_no_1" binding:"required,e164"`
	ContactNo2 string `json:"residential_contact_no_2" binding:"omitempty,e164"`
}
