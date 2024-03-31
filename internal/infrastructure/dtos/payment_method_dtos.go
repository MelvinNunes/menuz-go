package dtos

type CreatePaymentMethod struct {
	Name        string  `validator:"required" json:"name"`
	Key         string  `validator:"required" json:"key"`
	Type        string  `validator:"required" json:"type"`
	Currency    string  `validator:"required,currency_validator" json:"currency"`
	Description *string `json:"description"`
	Active      *bool   `json:"active"`
}
