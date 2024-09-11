package services

import (
	"context"
	"github.com/rodrigopero/coderhouse-challenge/src/domain"
	"github.com/rodrigopero/coderhouse-challenge/src/repositories"
)

type Account interface {
	GetAccount(ctx context.Context, username string) (*domain.Account, error)
}

type AccountDependencies struct {
	AccountRepository repositories.Account
}

type AccountImpl struct {
	AccountRepository repositories.Account
}

func NewAccountImpl(dependencies AccountDependencies) AccountImpl {
	return AccountImpl{
		AccountRepository: dependencies.AccountRepository,
	}
}

func (s AccountImpl) GetAccount(ctx context.Context, username string) (*domain.Account, error) {
	accountEntity, err := s.AccountRepository.GetAccountByUsername(ctx, username)

	if err != nil {
		return nil, err
	}

	return &domain.Account{Balance: accountEntity.Balance}, nil
}