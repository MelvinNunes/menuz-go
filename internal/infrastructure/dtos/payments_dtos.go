package dtos

import "github.com/google/uuid"

type PaymentMobileRequestDTO struct {
	PaymentMethod   string  `validate:"payment_method_validator" json:"payment_method"`
	PhoneNumberCode string  `validate:"phone_code_validator" json:"phone_number_code"`
	PhoneNumber     string  `validate:"required,min=9,max=9" json:"phone_number"` // Account Phone Number that will be money deducted
	Amount          float64 `validate:"required" json:"amount"`
	Description     *string `json:"description"`
}

type PaymentMobileDTO struct {
	UserID          uuid.UUID
	PaymentMethodID uint64
	SystemAmount    int64
	PhoneNumberCode string
	PhoneNumber     string
	Amount          float64
	Currency        string
	Description     *string
	TransactionType string
	Status          string
}
