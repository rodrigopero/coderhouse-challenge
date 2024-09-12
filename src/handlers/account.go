package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rodrigopero/coderhouse-challenge/src/handlers/dtos"
	"github.com/rodrigopero/coderhouse-challenge/src/services"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/api_error"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/auth"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/validation"
	"net/http"
	"strconv"
	"strings"
)

const (
	limitParam   = "limit"
	defaultLimit = 10
)

var (
	UnauthorizedUserError  = api_error.NewApiError(http.StatusUnauthorized, "user not authorized")
	InvalidLimitParamError = api_error.NewApiError(http.StatusBadRequest, "invalid limit param")
	InvalidBodyError       = api_error.NewApiError(http.StatusBadRequest, "invalid body")
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
		c.JSON(http.StatusUnauthorized, UnauthorizedUserError)
		return
	}

	accounts, err := h.AccountService.GetAllAccounts(ctx, username)
	if err != nil {
		c.JSON(api_error.GetStatus(err), err)
		return
	}

	var response []dtos.BalanceResponse

	for _, account := range accounts {
		response = append(response, dtos.BalanceResponse{Balance: account.Balance, Currency: account.Currency})
	}

	c.JSON(http.StatusOK, response)
}

func (h AccountImpl) Deposit(c *gin.Context) {
	ctx := c.Request.Context()

	username := auth_utils.GetAuthUser(c)
	if username == "" {
		c.JSON(http.StatusUnauthorized, UnauthorizedUserError)
		return
	}

	var dto dtos.DepositDTO
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, InvalidBodyError)
		return
	}
	err = validation.GetValidatorInstance().Struct(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, api_error.NewApiError(http.StatusBadRequest, validation.GetErrors(err.(validator.ValidationErrors))))
		return
	}

	err = validation.GetValidatorInstance().Var(dto.Currency, fmt.Sprintf("oneof=%s", strings.Join(allowedCurrencies, " ")))
	if err != nil {
		c.JSON(http.StatusBadRequest, InvalidCurrencyError)
		return
	}

	account, err := h.AccountService.Deposit(ctx, username, dto.Amount, dto.Currency)
	if err != nil {
		c.JSON(api_error.GetStatus(err), err)
	}

	c.JSON(http.StatusOK, dtos.BalanceResponse{Balance: account.Balance, Currency: account.Currency})
}

func (h AccountImpl) Withdraw(c *gin.Context) {
	ctx := c.Request.Context()

	username := auth_utils.GetAuthUser(c)
	if username == "" {
		c.JSON(http.StatusUnauthorized, UnauthorizedUserError)
		return
	}

	var dto dtos.WithdrawDTO
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, InvalidBodyError)
		return
	}
	err = validation.GetValidatorInstance().Struct(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, api_error.NewApiError(http.StatusBadRequest, validation.GetErrors(err.(validator.ValidationErrors))))
		return
	}

	err = validation.GetValidatorInstance().Var(dto.Currency, fmt.Sprintf("oneof=%s", strings.Join(allowedCurrencies, " ")))
	if err != nil {
		c.JSON(http.StatusBadRequest, InvalidCurrencyError)
		return
	}

	account, err := h.AccountService.Withdraw(ctx, username, dto.Amount, dto.Currency)
	if err != nil {
		c.JSON(api_error.GetStatus(err), err)
	}

	c.JSON(http.StatusOK, dtos.BalanceResponse{Balance: account.Balance, Currency: account.Currency})
}

func (h AccountImpl) GetTransactionHistory(c *gin.Context) {
	ctx := c.Request.Context()

	username := auth_utils.GetAuthUser(c)
	if username == "" {
		c.JSON(http.StatusUnauthorized, UnauthorizedUserError)
		return
	}

	limitStr := c.Query(limitParam)

	var limit int
	var err error
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, InvalidLimitParamError)
		}
	} else {
		limit = defaultLimit
	}

	transactions, err := h.AccountService.GetTransactionsHistory(ctx, username, limit)
	if err != nil {
		c.JSON(api_error.GetStatus(err), err)
		return
	}

	var response []dtos.TransactionResponse

	for _, transaction := range transactions {
		response = append(response, dtos.TransactionResponse{
			Amount:         transaction.Amount,
			Type:           transaction.Type,
			PartialBalance: transaction.PartialBalance,
			Date:           transaction.Date,
			Currency:       transaction.Currency,
		})
	}

	c.JSON(http.StatusOK, response)
}