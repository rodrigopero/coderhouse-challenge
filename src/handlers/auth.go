package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rodrigopero/coderhouse-challenge/src/handlers/dtos"
	"github.com/rodrigopero/coderhouse-challenge/src/services"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/api_error"
	auth_utils "github.com/rodrigopero/coderhouse-challenge/src/utils/auth"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/validation"
	"net/http"
)

const (
	authUserKey = "auth_user"
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

	err := c.BindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid body")
	}
	err = validation.GetValidatorInstance().Struct(dto)
	if err != nil {
		c.JSON(api_error.GetStatus(err), validation.GetErrorList(err.(validator.ValidationErrors)))
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
		token := c.Request.Header.Get("token")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "missing authorization token")
			return
		}

		valid := h.AuthService.IsValidToken(ctx, token)
		if !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, api_error.NewApiError(http.StatusUnauthorized, "unauthorized user"))
			return
		}

		username, err := h.AuthService.GetTokenUsername(ctx, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, api_error.NewApiError(http.StatusInternalServerError, "error validating token"))
			return
		}

		auth_utils.SetAuthUser(c, username)

		c.Next()
	}
}