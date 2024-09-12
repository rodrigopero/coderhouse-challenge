package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rodrigopero/coderhouse-challenge/src/handlers/dtos"
	"github.com/rodrigopero/coderhouse-challenge/src/services"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/api_error"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/validation"
	"net/http"
	"strings"
)

var (
	dollars = "USD"
	pesos   = "ARS"
	euros   = "EUR"

	allowedCurrencies = []string{dollars, pesos, euros}
)

type User interface {
	CreateUser(c *gin.Context)
}

type UserImpl struct {
	userService services.User
}

type UserDependencies struct {
	UserService services.User
}

func NewUserImpl(dependencies UserDependencies) UserImpl {
	return UserImpl{
		userService: dependencies.UserService,
	}
}

func (h UserImpl) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	var dto dtos.CreateUserDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, api_error.NewApiError(http.StatusBadRequest, "Invalid body"))
		return
	}
	err = validation.GetValidatorInstance().Struct(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, api_error.NewApiError(http.StatusBadRequest, validation.GetErrors(err.(validator.ValidationErrors))))
		return
	}

	for _, currency := range dto.Currencies {
		err = validation.GetValidatorInstance().Var(currency, fmt.Sprintf("oneof=%s", strings.Join(allowedCurrencies, " ")))
		if err != nil {
			c.JSON(http.StatusBadRequest, api_error.NewApiError(http.StatusBadRequest, fmt.Sprintf("The 'Currency' field only accepts the values %s", allowedCurrencies)))
			return
		}
	}

	err = h.userService.CreateUser(ctx, dto)
	if err != nil {
		c.JSON(api_error.GetStatus(err), err)
		return
	}

	c.Status(http.StatusCreated)
}
