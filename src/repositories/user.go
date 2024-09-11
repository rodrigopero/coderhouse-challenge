package repositories

import (
	"context"
	"database/sql"
	"errors"
	api_error "github.com/rodrigopero/coderhouse-challenge/src/utils/api_error"
	"net/http"
	"time"
)

const (
	insertStmt           = "INSERT INTO USERS (USERNAME, PASSWORD, STATUS, CREATION_DATE, MODIFICATION_DATE) VALUES (?,?,?,?,?)"
	selectByUsernameStmt = "SELECT ID, USERNAME, PASSWORD, STATUS, CREATION_DATE, MODIFICATION_DATE FROM USERS WHERE USERNAME = ?"
)

var (
	NotFoundError = api_error.NewApiError(http.StatusNotFound, "user not found")
)

type User interface {
	SaveUser(ctx context.Context, user UserEntity) error
	GetUserByUsername(ctx context.Context, username string) (*UserEntity, error)
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
	Id               int
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

func (r UserImpl) GetUserByUsername(ctx context.Context, username string) (*UserEntity, error) {
	row := r.database.QueryRowContext(ctx, selectByUsernameStmt, username)

	var user UserEntity
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Status, &user.CreationDate, &user.ModificationDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NotFoundError
		}
		return nil, err
	}

	return &user, nil
}