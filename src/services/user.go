package services

import (
	"context"
	"github.com/rodrigopero/coderhouse-challenge/src/handlers/dtos"
	"github.com/rodrigopero/coderhouse-challenge/src/repositories"
	"golang.org/x/crypto/bcrypt"
)

const (
	userActiveStatus = "active"
)

type User interface {
	CreateUser(ctx context.Context, user dtos.CreateUserDTO) error
}
type UserImpl struct {
	userRepository repositories.User
}

type UserDependencies struct {
	UserRepository repositories.User
}

func NewUserImpl(dependencies UserDependencies) UserImpl {
	return UserImpl{
		userRepository: dependencies.UserRepository,
	}
}

func (u UserImpl) CreateUser(ctx context.Context, dto dtos.CreateUserDTO) error {
	user := dto.ToDomain()

	hashedPass, err := hashPassword(dto.Password)
	if err != nil {
		return err
	}

	userEntity := repositories.UserEntity{
		Username: user.Username,
		Password: hashedPass,
		Status:   userActiveStatus,
	}

	return u.userRepository.SaveUser(ctx, userEntity)
}

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
