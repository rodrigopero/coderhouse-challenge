package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rodrigopero/coderhouse-challenge/src/utils/api_error"
	"net/http"
	"time"
)

const (
	insertUserStmt             = "INSERT INTO USERS (USERNAME, PASSWORD, CREATION_DATE, MODIFICATION_DATE, STATUS, LOGIN_ATTEMPTS) VALUES (?,?,?,?,?,?)"
	selectUserByUsernameStmt   = "SELECT ID, USERNAME, PASSWORD, CREATION_DATE, MODIFICATION_DATE, LOGIN_ATTEMPTS, STATUS FROM USERS WHERE USERNAME = ?"
	updateUserLoginAttemptStmt = "UPDATE USERS SET LOGIN_ATTEMPTS = ? WHERE ID = ?"
	updateUserStatusStmt       = "UPDATE USERS SET STATUS = ? WHERE ID = ?"
)

var (
	UserNotFoundError = api_error.NewApiError(http.StatusNotFound, "user not found")
)

type User interface {
	SaveUser(ctx context.Context, user UserEntity) (int, error)
	GetUserByUsername(ctx context.Context, username string) (*UserEntity, error)
	UpdateUserLoginAttempt(ctx context.Context, user UserEntity) error
	UpdateUserStatus(ctx context.Context, user UserEntity) error
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
	LoginAttempt     int
	Status           string
	Password         []byte
	CreationDate     string
	ModificationDate string
}

func (r UserImpl) SaveUser(ctx context.Context, user UserEntity) (int, error) {
	timeNow := time.Now()

	res, err := r.database.ExecContext(ctx, insertUserStmt, user.Username, user.Password, timeNow, timeNow, user.Status, user.LoginAttempt)
	if err != nil {
		return 0, err
	}

	insertedId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(insertedId), nil
}

func (r UserImpl) GetUserByUsername(ctx context.Context, username string) (*UserEntity, error) {
	row := r.database.QueryRowContext(ctx, selectUserByUsernameStmt, username)

	var user UserEntity
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.CreationDate, &user.ModificationDate, &user.LoginAttempt, &user.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, UserNotFoundError
		}
		return nil, err
	}

	return &user, nil
}

func (r UserImpl) UpdateUserLoginAttempt(ctx context.Context, user UserEntity) error {
	_, err := r.database.ExecContext(ctx, updateUserLoginAttemptStmt, user.LoginAttempt, user.Id)

	if err != nil {
		return err
	}

	return nil
}

func (r UserImpl) UpdateUserStatus(ctx context.Context, user UserEntity) error {
	_, err := r.database.ExecContext(ctx, updateUserStatusStmt, user.Status, user.Id)

	if err != nil {
		return err
	}

	return nil
}