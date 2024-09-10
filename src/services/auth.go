package services

type Auth interface {
}

type AuthDependencies struct{}

type AuthImpl struct {
}

func NewAuthImpl(dependencies AuthDependencies) AuthImpl {
	return AuthImpl{}
}
