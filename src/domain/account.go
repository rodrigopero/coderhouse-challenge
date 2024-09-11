package domain

type Account struct {
	Balance float64
}

type Transaction struct {
	AccountId      int
	UserId         int
	Amount         float64
	Type           string
	PartialBalance float64
}