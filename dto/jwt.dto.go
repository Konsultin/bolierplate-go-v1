package dto

type ValidateJwt_Payload struct {
	Subject  *string  `json:"subject,omitempty"`
	Audience []string `json:"audience,omitempty"`
}
