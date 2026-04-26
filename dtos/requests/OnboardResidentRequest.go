package requests

type OnboardResidentRequest struct {
	Name        string `json:"name" validate:"required,min=2"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Address     string `json:"address" validate:"required"`
}
