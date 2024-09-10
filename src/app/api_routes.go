package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigopero/coderhouse-challenge/src/app/dependencies"
	"github.com/rodrigopero/coderhouse-challenge/src/handlers"
	"net/http"
)

type RoutesDependencies struct {
	UserHandler handlers.User
}

func MapAPIRoutes(r *gin.Engine, depManager dependencies.Manager) {

	deps := RoutesDependencies{
		UserHandler: depManager.UserHandler(),
	}

	r.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "pong"}) })

	user := r.Group("/user")
	user.POST("", deps.UserHandler.CreateUser)
}
