package dtos

import (
	"time"
)

type DepositDTO struct {
	Amount   float64 `json:"amount" validate:"required,gt=0"`
	Currency string  `json:"currency" validate:"required"`
}

type WithdrawDTO struct {
	Amount   float64 `json:"amount" validate:"required,gt=0"`
	Currency string  `json:"currency" validate:"required"`
}

type BalanceResponse struct {
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

type TransactionResponse struct {
	Amount         float64   `json:"amount"`
	Type           string    `json:"type"`
	PartialBalance float64   `json:"partial_balance"`
	Date           time.Time `json:"date"`
	Currency       string    `json:"currency"`
}