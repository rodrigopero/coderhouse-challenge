package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rodrigopero/coderhouse-challenge/src/handlers/dtos"
	"github.com/rodrigopero/coderhouse-challenge/src/services"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/api_error"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/auth"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/validation"
	"net/http"
	"strconv"
)

const (
	limitParam   = "limit"
	defaultLimit = 10
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

	c.JSON(http.StatusOK, dtos.BalanceResponse{Balance: account.Balance})
}

func (h AccountImpl) Deposit(c *gin.Context) {
	ctx := c.Request.Context()

	username := auth_utils.GetAuthUser(c)
	if username == "" {
		c.JSON(http.StatusUnauthorized, api_error.NewApiError(http.StatusUnauthorized, "unauthorized user"))
		return
	}

	var dto dtos.DepositDTO

	err := c.BindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	err = validation.GetValidatorInstance().Struct(dto)
	if err != nil {
		c.JSON(api_error.GetStatus(err), validation.GetErrorList(err.(validator.ValidationErrors)))
		return
	}

	account, err := h.AccountService.Deposit(ctx, username, dto.Amount)
	if err != nil {
		c.JSON(api_error.GetStatus(err), err)
	}

	c.JSON(http.StatusOK, dtos.BalanceResponse{Balance: account.Balance})

}

func (h AccountImpl) Withdraw(c *gin.Context) {
	ctx := c.Request.Context()

	username := auth_utils.GetAuthUser(c)
	if username == "" {
		c.JSON(http.StatusUnauthorized, api_error.NewApiError(http.StatusUnauthorized, "unauthorized user"))
		return
	}

	var dto dtos.WithdrawDTO

	err := c.BindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	err = validation.GetValidatorInstance().Struct(dto)
	if err != nil {
		c.JSON(api_error.GetStatus(err), validation.GetErrorList(err.(validator.ValidationErrors)))
		return
	}

	account, err := h.AccountService.Withdraw(ctx, username, dto.Amount)
	if err != nil {
		c.JSON(api_error.GetStatus(err), err)
	}

	c.JSON(http.StatusOK, dtos.BalanceResponse{Balance: account.Balance})

}

func (h AccountImpl) GetTransactionHistory(c *gin.Context) {
	ctx := c.Request.Context()

	username := auth_utils.GetAuthUser(c)
	if username == "" {
		c.JSON(http.StatusUnauthorized, api_error.NewApiError(http.StatusUnauthorized, "unauthorized user"))
		return
	}

	limitStr := c.Query("limit")

	var limit int
	var err error
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, api_error.NewApiError(http.StatusBadRequest, "invalid limit param"))
		}
	} else {
		limit = defaultLimit
	}

	transactions, err := h.AccountService.GetTransactionsHistory(ctx, username, limit)
	if err != nil {
		c.JSON(api_error.GetStatus(err), err)
		return
	}

	c.JSON(http.StatusOK, transactions)
}