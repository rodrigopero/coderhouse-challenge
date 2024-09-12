package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rodrigopero/coderhouse-challenge/src/handlers/dtos"
	"github.com/rodrigopero/coderhouse-challenge/src/services"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/api_error"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/validation"
	"net/http"
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

	err = h.userService.CreateUser(ctx, dto)
	if err != nil {
		c.JSON(api_error.GetStatus(err), err)
		return
	}

	c.Status(http.StatusCreated)
}
