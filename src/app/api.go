package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigopero/coderhouse-challenge/src/app/dependencies"
	"github.com/rodrigopero/coderhouse-challenge/src/repositories/clients"
	"net/http"
)

const (
	dbPath = "./db/"
)

func Run() error {
	err := InitializeDB()
	if err != nil {
		return err
	}

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

func InitializeDB() error {
	_, err := clients.InitializeDbSqlite(dbPath)
	return err
}