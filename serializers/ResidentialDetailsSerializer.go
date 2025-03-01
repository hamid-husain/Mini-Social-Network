package serializers

type ResidentialDetailsInput struct {
	Address    string `json:"address" binding:"required"`
	City       string `json:"city" binding:"required"`
	State      string `json:"state" binding:"required"`
	Country    string `json:"country" binding:"required"`
	ContactNo1 string `json:"contact_no_1" binding:"required,e164"`
	ContactNo2 string `json:"contact_no_2" binding:"omitempty,e164"`
}

type ResidentialDetailsResponse struct {
	ID         uint   `json:"id"`
	UserID     uint   `json:"user_id"`
	Address    string `json:"address" binding:"required"`
	City       string `json:"city" binding:"required"`
	State      string `json:"state" binding:"required"`
	Country    string `json:"country" binding:"required"`
	ContactNo1 string `json:"contact_no_1" binding:"required,e164"`
	ContactNo2 string `json:"contact_no_2" binding:"omitempty,e164"`
}
