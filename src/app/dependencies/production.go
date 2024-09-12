package dependencies

import (
	"database/sql"
	"github.com/rodrigopero/coderhouse-challenge/src/handlers"
	"github.com/rodrigopero/coderhouse-challenge/src/repositories"
	"github.com/rodrigopero/coderhouse-challenge/src/repositories/clients"
	"github.com/rodrigopero/coderhouse-challenge/src/services"
	"log"
)

type Production struct {
	userHandler           handlers.User
	authHandler           handlers.Auth
	accountHandler        handlers.Account
	userService           services.User
	authService           services.Auth
	accountService        services.Account
	userRepository        repositories.User
	accountRepository     repositories.Account
	transactionRepository repositories.Transaction
	databaseClient        *sql.DB
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

func (p *Production) AccountHandler() handlers.Account {
	if p.accountHandler == nil {
		dependencies := handlers.AccountDependencies{
			AccountService: p.AccountService(),
		}
		p.accountHandler = handlers.NewAccountImpl(dependencies)
	}
	return p.accountHandler
}

// Services
func (p *Production) UserService() services.User {
	if p.userService == nil {
		dependencies := services.UserDependencies{
			UserRepository:    p.UserRepository(),
			AccountRepository: p.AccountRepository(),
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

func (p *Production) AccountService() services.Account {
	if p.accountService == nil {
		dependencies := services.AccountDependencies{
			AccountRepository:     p.AccountRepository(),
			TransactionRepository: p.TransactionRepository(),
		}
		p.accountService = services.NewAccountImpl(dependencies)
	}
	return p.accountService
}

// Repositories
func (p *Production) UserRepository() repositories.User {
	if p.userRepository == nil {
		dependencies := repositories.UserDependencies{
			Database: p.DatabaseClient(),
		}
		p.userRepository = repositories.NewUserImpl(dependencies)
	}
	return p.userRepository
}

func (p *Production) AccountRepository() repositories.Account {
	if p.accountRepository == nil {
		dependencies := repositories.AccountDependencies{
			Database: p.DatabaseClient(),
		}
		p.accountRepository = repositories.NewAccountImpl(dependencies)
	}
	return p.accountRepository
}

func (p *Production) TransactionRepository() repositories.Transaction {
	if p.transactionRepository == nil {
		dependencies := repositories.TransactionDependencies{
			Database: p.DatabaseClient(),
		}
		p.transactionRepository = repositories.NewTransactionImpl(dependencies)
	}
	return p.transactionRepository
}

// Others
func (p *Production) DatabaseClient() *sql.DB {
	if p.databaseClient == nil {
		client, err := clients.NewSQLite()
		if err != nil {
			log.Panicf("error initializing database: %s", err.Error())
		}

		p.databaseClient = client
	}
	return p.databaseClient
}
