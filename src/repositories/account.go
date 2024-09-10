package repositories

type Account interface {
}

type AccountDependencies struct{}

type AccountImpl struct {
}

func NewAccountImpl(dependencies AccountDependencies) AccountImpl {
	return AccountImpl{}
}
