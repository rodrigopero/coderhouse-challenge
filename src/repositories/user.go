package repositories

import (
	"context"
	"database/sql"
	"time"
)

const (
	insertStmt = "INSERT INTO USERS (username, password, status, creation_date, modification_date) VALUES (?,?,?,?,?)"
)

type User interface {
	SaveUser(ctx context.Context, user UserEntity) error
}

type UserImpl struct {
	database *sql.DB
}

type UserDependencies struct {
	Database *sql.DB
}

func NewUserImpl(dependencies UserDependencies) UserImpl {
	return UserImpl{
		database: dependencies.Database,
	}
}

type UserEntity struct {
	Username         string
	Password         []byte
	Status           string
	CreationDate     string
	ModificationDate string
}

func (r UserImpl) SaveUser(ctx context.Context, user UserEntity) error {

	timeNow := time.Now()

	_, err := r.database.ExecContext(ctx, insertStmt, user.Username, user.Password, user.Status, timeNow, timeNow)

	if err != nil {
		return err
	}

	return nil
}
