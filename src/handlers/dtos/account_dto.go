package dtos

import "time"

type DepositDTO struct {
	Amount float64 `json:"amount" validate:"required,gt=0"`
}

type WithdrawDTO struct {
	Amount float64 `json:"amount" validate:"required,gt=0"`
}

type BalanceResponse struct {
	Amount float64 `json:"amount"`
}

type HistoryResponse struct {
	Movements []TransactionResponse `json:"movements"`
}

type TransactionResponse struct {
	Amount         float64   `json:"amount"`
	Type           string    `json:"type"`
	PartialBalance float64   `json:"partial_balance"`
	Date           time.Time `json:"date"`
}