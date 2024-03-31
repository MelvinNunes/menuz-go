package dtos

type TransactionCreateDTO struct {
	Amount            int     `json:"amount"`
	Currency          string  `json:"currency"`
	Description       *string `json:"description"`
	PushedPhoneNumber *string `json:"pushed_phone_number"`
	UserID            string  `json:"user_id"`
	PaymentMethodID   uint64  `json:"payment_method_id"`
	TransactionType   string  `json:"transaction_type"`
	Status            string  `json:"status"`
}

type TransactionVM struct {
	ID                string `json:"id"`
	Amount            int    `json:"amount"`
	FeeAmount         int    `json:"fee_amount"`
	BrandName         string `json:"brand_name"`
	Currency          string `json:"currency"`
	PaymentMethod     string `json:"payment_method"`
	PushedPhoneNumber string `json:"pushed_phone_number"`
	TransactionType   string `json:"transaction_type"`
	Sender            UserVM `json:"sender"`
	Receiver          UserVM `json:"receiver"`
	Description       string `json:"description"`
	Voucher           string `json:"voucher"`
	Category          string `json:"category"`
	Status            string `json:"status"`
	CreatedAt         string `json:"created_at"`
}

type TransactionVM2 struct {
	ID                string `json:"id"`
	Amount            int    `json:"amount"`
	Currency          string `json:"currency"`
	PaymentMethod     string `json:"payment_method"`
	PushedPhoneNumber string `json:"pushed_phone_number"`
	TransactionType   string `json:"transaction_type"`
	Type              string `json:"type"`
	Message           string `json:"message"`
	BrandName         string `json:"brand_name"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	Voucher           string `json:"voucher"`
	Category          string `json:"category"`
	Avatar            string `json:"avatar"`
	CreatedAt         string `json:"created_at"`
}
