package dependencies

import (
	"github.com/rodrigopero/coderhouse-challenge/src/handlers"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/env"
	"log"
)

const (
	environmentEnv        = "ENVIRONMENT"
	productionEnvironment = "production"
)

type Manager interface {
	UserHandler() handlers.User
	AuthHandler() handlers.Auth
	AccountHandler() handlers.Account
}

func NewManager() Manager {
	environment := env.GetEnvVar(environmentEnv)

	switch environment {
	case productionEnvironment:
		return &Production{}
	default:
		log.Fatal("error loading environment settings")
		return nil
	}
}
