package services

import (
	"context"
	"errors"
	"github.com/rodrigopero/coderhouse-challenge/src/handlers/dtos"
	"github.com/rodrigopero/coderhouse-challenge/src/repositories"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/api_error"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const (
	activeUserStatus  = "active"
	blockedUserStatus = "blocked"
)

var (
	AlreadyExistsErr = api_error.NewApiError(http.StatusBadRequest, "user already exists")
)

type User interface {
	CreateUser(ctx context.Context, user dtos.CreateUserDTO) error
}
type UserImpl struct {
	userRepository    repositories.User
	accountRepository repositories.Account
}

type UserDependencies struct {
	UserRepository    repositories.User
	AccountRepository repositories.Account
}

func NewUserImpl(dependencies UserDependencies) UserImpl {
	return UserImpl{
		userRepository:    dependencies.UserRepository,
		accountRepository: dependencies.AccountRepository,
	}
}

func (s UserImpl) CreateUser(ctx context.Context, dto dtos.CreateUserDTO) error {
	user := dto.ToDomain()

	_, err := s.userRepository.GetUserByUsername(ctx, user.Username)
	if err == nil {
		return AlreadyExistsErr
	} else if !errors.Is(err, repositories.UserNotFoundError) {
		return err
	}

	hashedPass, err := hashPassword(dto.Password)
	if err != nil {
		return err
	}

	userEntity := repositories.UserEntity{
		Username: user.Username,
		Password: hashedPass,
		Status:   activeUserStatus,
	}

	userId, err := s.userRepository.SaveUser(ctx, userEntity)
	if err != nil {
		return err
	}

	for _, currency := range dto.Currencies {
		accountEntity := repositories.AccountEntity{
			UserId:   userId,
			Currency: currency,
		}
		err = s.accountRepository.SaveAccount(ctx, accountEntity)
		if err != nil {
			return err
		}
	}

	return nil
}

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}