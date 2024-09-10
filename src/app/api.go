package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigopero/coderhouse-challenge/src/app/dependencies"
)

func Run() error {
	r := gin.Default()

	MapAPIRoutes(r, dependencies.NewManager())

	return r.Run()
}
