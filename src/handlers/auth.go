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
)

const (
	tokenHeader = "token"
)

var (
	UnexpectedError   = api_error.NewApiError(http.StatusUnauthorized, "unexpected error")
	MissingTokenError = api_error.NewApiError(http.StatusForbidden, "missing authorization token")
)

type Auth interface {
	Authenticate(c *gin.Context)
	AuthMiddleware() gin.HandlerFunc
}

type AuthDependencies struct {
	AuthService services.Auth
}

type AuthImpl struct {
	AuthService services.Auth
}

func NewAuthImpl(dependencies AuthDependencies) AuthImpl {
	return AuthImpl{
		AuthService: dependencies.AuthService,
	}
}

func (h AuthImpl) Authenticate(c *gin.Context) {
	ctx := c.Request.Context()
	var dto dtos.AuthorizationDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, InvalidBodyError)
	}

	err = validation.GetValidatorInstance().Struct(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, api_error.NewApiError(http.StatusBadRequest, validation.GetErrors(err.(validator.ValidationErrors))))
		return
	}

	token, err := h.AuthService.AuthenticateUser(ctx, dto)
	if err != nil {
		c.JSON(api_error.GetStatus(err), err)
		return
	}

	c.JSON(http.StatusOK, dtos.AuthorizationResponse{Token: token})
}

func (h AuthImpl) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		token := c.Request.Header.Get(tokenHeader)

		if token == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, MissingTokenError)
			return
		}

		valid := h.AuthService.IsValidToken(ctx, token)
		if !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedUserError)
			return
		}

		username, err := h.AuthService.GetTokenUsername(ctx, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, UnexpectedError)
			return
		}

		auth_utils.SetAuthUser(c, username)

		c.Next()
	}
}