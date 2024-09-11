package dependencies

import (
	"github.com/rodrigopero/coderhouse-challenge/src/handlers"
	"github.com/rodrigopero/coderhouse-challenge/src/repositories"
	"github.com/rodrigopero/coderhouse-challenge/src/repositories/clients"
	"github.com/rodrigopero/coderhouse-challenge/src/services"
	"log"
)

type Production struct {
	userHandler    handlers.User
	authHandler    handlers.Auth
	userService    services.User
	authService    services.Auth
	userRepository repositories.User
}

// Handlers
func (p *Production) UserHandler() handlers.User {
	if p.userHandler == nil {
		dependencies := handlers.UserDependencies{
			UserService: p.UserService(),
		}
		p.userHandler = handlers.NewUserImpl(dependencies)
	}
	return p.userHandler
}

func (p *Production) AuthHandler() handlers.Auth {
	if p.authHandler == nil {
		dependencies := handlers.AuthDependencies{
			AuthService: p.AuthService(),
		}
		p.authHandler = handlers.NewAuthImpl(dependencies)
	}
	return p.authHandler
}

// Services
func (p *Production) UserService() services.User {
	if p.userService == nil {
		dependencies := services.UserDependencies{
			UserRepository: p.UserRepository(),
		}
		p.userService = services.NewUserImpl(dependencies)
	}
	return p.userService
}

func (p *Production) AuthService() services.Auth {
	if p.authService == nil {
		dependencies := services.AuthDependencies{
			UserRepository: p.UserRepository(),
		}
		p.authService = services.NewAuthImpl(dependencies)
	}
	return p.authService
}

// Repositories
func (p *Production) UserRepository() repositories.User {
	if p.userRepository == nil {
		client, err := clients.NewSQLite()
		if err != nil {
			log.Panicf("error initializing database: %s", err.Error())
		}

		dependencies := repositories.UserDependencies{
			Database: client,
		}
		p.userRepository = repositories.NewUserImpl(dependencies)
	}
	return p.userRepository
}