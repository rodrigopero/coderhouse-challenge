package services

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rodrigopero/coderhouse-challenge/src/domain"
	"github.com/rodrigopero/coderhouse-challenge/src/handlers/dtos"
	"github.com/rodrigopero/coderhouse-challenge/src/repositories"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/api_error"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

const (
	jwtKey            = "c0D3rH0u5E-Ch411eNg3"
	jwtExpirationTime = time.Hour * time.Duration(15)
	maxLoginAttempts  = 3
)

var (
	UnauthorizedError            = api_error.NewApiError(http.StatusUnauthorized, "user not authorized")
	IncorrectAuthenticationError = api_error.NewApiError(http.StatusUnauthorized, "incorrect username or password")
	BlockedUserError             = api_error.NewApiError(http.StatusUnauthorized, "user is blocked")
)

type Auth interface {
	AuthenticateUser(ctx context.Context, dto dtos.AuthorizationDTO) (string, error)
	IsValidToken(ctx context.Context, token string) bool
	GetTokenUsername(ctx context.Context, token string) (string, error)
}

type AuthDependencies struct {
	UserRepository repositories.User
}

type AuthImpl struct {
	UserRepository repositories.User
}

func NewAuthImpl(dependencies AuthDependencies) AuthImpl {
	return AuthImpl{
		UserRepository: dependencies.UserRepository,
	}
}

func (s AuthImpl) AuthenticateUser(ctx context.Context, dto dtos.AuthorizationDTO) (string, error) {

	user, err := s.UserRepository.GetUserByUsername(ctx, dto.Username)
	if errors.Is(err, repositories.UserNotFoundError) {
		return "", UnauthorizedError
	}

	if user.Status == blockedUserStatus {
		return "", BlockedUserError
	}

	passwordValid := checkValidPassword(user.Password, dto.Password)
	if !passwordValid {
		user.LoginAttempt += 1
		err = s.UserRepository.UpdateUserLoginAttempt(ctx, *user)
		if err != nil {
			return "", err
		}
		if user.LoginAttempt > maxLoginAttempts {
			user.Status = blockedUserStatus
			err = s.UserRepository.UpdateUserStatus(ctx, *user)
			return "", BlockedUserError
		}

		return "", IncorrectAuthenticationError
	}

	if user.LoginAttempt > 0 {
		user.LoginAttempt = 0
		err = s.UserRepository.UpdateUserLoginAttempt(ctx, *user)
		if err != nil {
			return "", err
		}
	}

	token, err := generateToken(user.Username)
	if err != nil {
		return "", api_error.NewApiError(http.StatusInternalServerError, "error generating authorization token")
	}

	return token, nil
}

func (s AuthImpl) IsValidToken(ctx context.Context, token string) bool {
	tokenData, _ := jwt.ParseWithClaims(token, &domain.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid method")
		}

		return []byte(jwtKey), nil
	})

	if _, ok := tokenData.Claims.(*domain.CustomClaims); ok && tokenData.Valid {
		return true
	} else {
		return false
	}
}

func (s AuthImpl) GetTokenUsername(ctx context.Context, asd string) (string, error) {
	token, err := jwt.ParseWithClaims(asd, &domain.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, _ := token.Claims.(*domain.CustomClaims)
	return claims.Username, nil
}

func checkValidPassword(userPassword []byte, authPassword string) bool {
	err := bcrypt.CompareHashAndPassword(userPassword, []byte(authPassword))
	return err == nil
}

func generateToken(username string) (string, error) {
	claims := domain.CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtExpirationTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}