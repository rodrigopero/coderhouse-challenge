package handlers

import "github.com/gin-gonic/gin"

type Auth interface {
	Authenticate(c *gin.Context)
}

type AuthDependencies struct{}

type AuthImpl struct {
}

func NewAuthImpl(dependencies AuthDependencies) AuthImpl {
	return AuthImpl{}
}

func (h AuthImpl) Authenticate(c *gin.Context) {}
