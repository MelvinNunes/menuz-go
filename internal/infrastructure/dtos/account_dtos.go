package dtos

import "time"

type LoginDTO struct {
	Email    string `validate:"required,email" json:"email" `
	Password string `validate:"required" json:"password"`
}

type CreateAccount struct {
	Email           string    `validate:"required,email" json:"email" `
	Password        string    `validate:"required,min=5" json:"password"`
	FirstName       string    `validate:"required" json:"first_name"`
	LastName        string    `validate:"required" json:"last_name"`
	PhoneNumberCode string    `validate:"phone_code_validator" json:"phone_number_code"`
	PhoneNumber     string    `validate:"required" json:"phone_number"`
	AppLanguage     string    `validate:"lang_validator" json:"app_language"`
	Gender          string    `validate:"required" json:"gender"`
	DateOfBirth     time.Time `validate:"required" json:"date_of_birth"`
	Role            string    `validate:"role_validator" json:"role"`
}
