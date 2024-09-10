package repositories

type Auth interface {
}

type AuthDependencies struct{}

type AuthSQLiteImpl struct{}

func NewAuthSQLiteRepository() AuthSQLiteImpl {
	return AuthSQLiteImpl{}
}
