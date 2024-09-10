package dtos

import "github.com/rodrigopero/coderhouse-challenge/src/domain"

type CreateUserDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (d CreateUserDTO) ToDomain() domain.User {
	return domain.User{
		Username: d.Username,
		Password: d.Password,
	}
}
