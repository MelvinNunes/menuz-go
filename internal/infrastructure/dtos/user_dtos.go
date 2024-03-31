package dtos

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserDTO struct {
	VerificationID             uint64 `json:"verification_id"`
	PhoneNumberCode            string `json:"phone_number_code"`
	PhoneNumber                string `json:"phone_number"`
	Role                       string `json:"role"`
	DeviceType                 string `json:"device_type"`
	DeviceFirebaseToken        string `json:"device_firebase_token"`
	IdentifyStatus             string `json:"identify_status"`
	AppLanguage                string `json:"app_language"`
	Category                   string `json:"category"`
	CustomerRegistrationStatus string `json:"customer_registration_status"`
}

type UpdateUserDTO struct {
	DeviceType                 string    `json:"device_type"`
	DeviceFirebaseToken        string    `json:"device_firebase_token"`
	Username                   string    `json:"username"`
	IdentityNumber             string    `json:"identity_number"`
	Password                   string    `json:"password"`
	ActivatedAt                time.Time `json:"activated_at"`
	IdentifyStatus             string    `json:"identify_status"`
	Category                   string    `json:"category"`
	CustomerRegistrationStatus string    `json:"customer_registration_status"`
	MerchantRole               string    `json:"merchant_role"`
	MerchantNationality        string    `json:"merchant_nationality"`
}

type UserVM struct {
	ID           uuid.UUID `json:"id"`
	Code         string    `json:"code"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	AccountTypes []string  `json:"account_types"`
	CreatedAt    string    `json:"created_at"`
}
