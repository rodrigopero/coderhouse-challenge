package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigopero/coderhouse-challenge/src/app/dependencies"
	"github.com/rodrigopero/coderhouse-challenge/src/handlers"
)

type RoutesDependencies struct {
	UserHandler    handlers.User
	AuthHandler    handlers.Auth
	AccountHandler handlers.Account
}

func MapAPIRoutes(r *gin.Engine, depManager dependencies.Manager) {

	deps := RoutesDependencies{
		UserHandler:    depManager.UserHandler(),
		AuthHandler:    depManager.AuthHandler(),
		AccountHandler: depManager.AccountHandler(),
	}

	r.POST("/authorize", deps.AuthHandler.Authenticate)

	user := r.Group("/user")
	user.POST("", deps.UserHandler.CreateUser)

	account := r.Group("/account").Use(deps.AuthHandler.AuthMiddleware())
	account.GET("/balance", deps.AccountHandler.GetBalance)
	account.POST("/deposit", deps.AccountHandler.Deposit)
	account.POST("/withdraw", deps.AccountHandler.Withdraw)
	account.POST("/transactions", deps.AccountHandler.GetTransactionHistory)

}