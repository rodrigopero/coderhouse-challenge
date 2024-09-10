package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigopero/coderhouse-challenge/src/handlers/dtos"
	"github.com/rodrigopero/coderhouse-challenge/src/services"
	"github.com/rodrigopero/coderhouse-challenge/src/utils"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/errors"
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
	}
	err = utils.GetValidatorInstance().Struct(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	err = h.userService.CreateUser(ctx, dto)
	if err != nil {
		c.JSON(errors.GetStatus(err), errors.GetMessage(err))
		return
	}

	c.Status(http.StatusCreated)
}
