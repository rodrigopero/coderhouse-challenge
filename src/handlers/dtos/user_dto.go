package dtos

import "github.com/rodrigopero/coderhouse-challenge/src/domain"

type CreateUserDTO struct {
	Username string `json:"username" validate:"required,alphanum,gte=8,lte=32"`
	Password string `json:"password" validate:"required,gte=8,lte=64"`
}

func (d CreateUserDTO) ToDomain() domain.User {
	return domain.User{
		Username: d.Username,
		Password: d.Password,
	}
}
