package handlers

type Transaction interface {
}

type TransactionDependencies struct{}

type TransactionImpl struct {
}

func NewTransactionImpl(dependencies TransactionDependencies) TransactionImpl {
	return TransactionImpl{}
}
