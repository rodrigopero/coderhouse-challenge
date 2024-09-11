package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigopero/coderhouse-challenge/src/handlers/dtos"
	"github.com/rodrigopero/coderhouse-challenge/src/services"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/api_error"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/auth"
	"net/http"
)

type Account interface {
	GetBalance(c *gin.Context)
	Deposit(c *gin.Context)
	Withdraw(c *gin.Context)
	GetTransactionHistory(c *gin.Context)
}

type AccountDependencies struct {
	AccountService services.Account
}

type AccountImpl struct {
	AccountService services.Account
}

func NewAccountImpl(dependencies AccountDependencies) AccountImpl {
	return AccountImpl{
		AccountService: dependencies.AccountService,
	}
}

func (h AccountImpl) GetBalance(c *gin.Context) {
	ctx := c.Request.Context()
	username := auth_utils.GetAuthUser(c)

	if username == "" {
		c.JSON(http.StatusUnauthorized, api_error.NewApiError(http.StatusUnauthorized, "unauthorized user"))
		return
	}

	account, err := h.AccountService.GetAccount(ctx, username)
	if err != nil {
		c.JSON(api_error.GetStatus(err), err)
		return
	}

	c.JSON(http.StatusOK, dtos.BalanceResponse{Amount: account.Balance})
}

func (t AccountImpl) Deposit(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t AccountImpl) Withdraw(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t AccountImpl) GetTransactionHistory(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}