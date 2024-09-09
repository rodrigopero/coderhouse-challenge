package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigopero/coderhouse-challenge/src/app/dependencies"
	"net/http"
)

func Run() error {
	//initializeDB

	return startAPI()

}

func startAPI() error {
	r := gin.Default()

	MapAPIRoutes(r, dependencies.NewManager())

	return r.Run()
}

func MapAPIRoutes(r *gin.Engine, manager dependencies.Manager) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}