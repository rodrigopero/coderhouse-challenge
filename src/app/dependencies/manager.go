package dependencies

import "github.com/rodrigopero/coderhouse-challenge/src/handlers"

const (
	productionEnvironment = "production"
)

type Manager interface {
	UserHandler() handlers.User
}

func NewManager() Manager {
	return &Production{}
}
