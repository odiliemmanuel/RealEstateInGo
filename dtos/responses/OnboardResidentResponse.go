package responses

type OnboardResidentResponse struct {
	ResidentID     string `json:"residentId"`
	ResidentName   string `json:"residentName"`
	DateRegistered string `json:"dateRegistered"`
}
