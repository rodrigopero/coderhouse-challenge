package domain

import "github.com/golang-jwt/jwt/v5"

type Authorization struct {
	Username string
	Password string
}

type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
