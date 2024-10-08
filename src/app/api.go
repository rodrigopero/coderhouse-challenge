package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigopero/coderhouse-challenge/src/app/dependencies"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/env"
)

func Run() error {
	env.SetEnvVars()

	r := gin.Default()
	MapAPIRoutes(r, dependencies.NewManager())

	return r.Run()
}
