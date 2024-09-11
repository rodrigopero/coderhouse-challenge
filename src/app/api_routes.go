package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigopero/coderhouse-challenge/src/app/dependencies"
	"github.com/rodrigopero/coderhouse-challenge/src/handlers"
	"net/http"
)

type RoutesDependencies struct {
	UserHandler handlers.User
	AuthHandler handlers.Auth
}

func MapAPIRoutes(r *gin.Engine, depManager dependencies.Manager) {

	deps := RoutesDependencies{
		UserHandler: depManager.UserHandler(),
		AuthHandler: depManager.AuthHandler(),
	}

	r.POST("/authorize", deps.AuthHandler.Authenticate)

	user := r.Group("/user").Use(deps.AuthHandler.AuthMiddleware())
	user.POST("", deps.UserHandler.CreateUser)

	user.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "pong"}) })

}