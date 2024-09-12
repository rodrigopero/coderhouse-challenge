package services

import (
	"context"
	"github.com/rodrigopero/coderhouse-challenge/src/domain"
	"github.com/rodrigopero/coderhouse-challenge/src/repositories"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/api_error"
	"net/http"
	"time"
)

const (
	DepositType  = "deposit"
	WithdrawType = "withdraw"
)

type Account interface {
	GetAccount(ctx context.Context, username string) (*domain.Account, error)
	Deposit(ctx context.Context, username string, amount float64) (*domain.Account, error)
	Withdraw(ctx context.Context, username string, amount float64) (*domain.Account, error)
	GetTransactionsHistory(ctx context.Context, username string, limit int) ([]domain.Transaction, error)
}

type AccountDependencies struct {
	AccountRepository     repositories.Account
	TransactionRepository repositories.Transaction
}

type AccountImpl struct {
	AccountRepository     repositories.Account
	TransactionRepository repositories.Transaction
}

func NewAccountImpl(dependencies AccountDependencies) AccountImpl {
	return AccountImpl{
		AccountRepository:     dependencies.AccountRepository,
		TransactionRepository: dependencies.TransactionRepository,
	}
}

func (s AccountImpl) GetAccount(ctx context.Context, username string) (*domain.Account, error) {
	accountEntity, err := s.AccountRepository.GetAccountByUsername(ctx, username)

	if err != nil {
		return nil, err
	}

	return &domain.Account{Balance: accountEntity.Balance}, nil
}

func (s AccountImpl) Deposit(ctx context.Context, username string, amount float64) (*domain.Account, error) {
	accountEntity, err := s.AccountRepository.GetAccountByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	accountEntity.Balance += amount
	err = s.AccountRepository.UpdateAccountBalance(ctx, *accountEntity)
	if err != nil {
		return nil, err
	}

	transactionEntity := repositories.TransactionEntity{
		UserId:         accountEntity.UserId,
		AccountId:      accountEntity.Id,
		Amount:         amount,
		PartialBalance: accountEntity.Balance,
		Type:           DepositType,
	}

	err = s.TransactionRepository.SaveTransaction(ctx, transactionEntity)
	if err != nil {
		return nil, err
	}

	return &domain.Account{Balance: accountEntity.Balance}, nil
}

func (s AccountImpl) Withdraw(ctx context.Context, username string, amount float64) (*domain.Account, error) {
	accountEntity, err := s.AccountRepository.GetAccountByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if accountEntity.Balance-amount < 0 {
		return nil, api_error.NewApiError(http.StatusBadRequest, "the account has insufficient funds")
	}

	accountEntity.Balance -= amount
	err = s.AccountRepository.UpdateAccountBalance(ctx, *accountEntity)
	if err != nil {
		return nil, err
	}

	transactionEntity := repositories.TransactionEntity{
		UserId:         accountEntity.UserId,
		AccountId:      accountEntity.Id,
		Amount:         amount,
		PartialBalance: accountEntity.Balance,
		Type:           WithdrawType,
	}

	err = s.TransactionRepository.SaveTransaction(ctx, transactionEntity)
	if err != nil {
		return nil, err
	}

	return &domain.Account{Balance: accountEntity.Balance}, nil
}

func (s AccountImpl) GetTransactionsHistory(ctx context.Context, username string, limit int) ([]domain.Transaction, error) {
	accountEntity, err := s.AccountRepository.GetAccountByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	transactionEntityList, err := s.TransactionRepository.GetTransactionsWithLimit(ctx, username, accountEntity.Id, limit)
	if err != nil {
		return nil, err
	}

	var transactions []domain.Transaction

	for _, transactionEntity := range transactionEntityList {

		parsedTime, err := parseTime(transactionEntity.Date)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions,
			domain.Transaction{
				AccountId:      transactionEntity.AccountId,
				UserId:         transactionEntity.UserId,
				Amount:         transactionEntity.Amount,
				Type:           transactionEntity.Type,
				PartialBalance: transactionEntity.PartialBalance,
				Date:           parsedTime,
			},
		)
	}

	return transactions, nil
}

func parseTime(date string) (time.Time, error) {
	layout := "2006-01-02 15:04:05.999999999-07:00"

	parsedTime, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
