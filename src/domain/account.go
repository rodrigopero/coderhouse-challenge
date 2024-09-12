package domain

import (
	"time"
)

type Account struct {
	Balance  float64
	Currency string
}

type Transaction struct {
	AccountId      int
	UserId         int
	Amount         float64
	Type           string
	PartialBalance float64
	Date           time.Time
	Currency       string
}
