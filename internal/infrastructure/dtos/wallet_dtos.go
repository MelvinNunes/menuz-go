package dtos

import "github.com/google/uuid"

type CreateWallet struct {
	UserID          uuid.UUID `json:"user_id"`
	Currency        string    `json:"currency"`
	Balance         int64     `json:"balance"`
	TopUpBalance    int64     `json:"top_up_balance"`
	WithdrawBalance int64     `json:"withdraw_balance"`
	Decimals        int       `json:"decimals"`
}
