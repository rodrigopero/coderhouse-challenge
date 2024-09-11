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
	initialAccountBalance = 0.0

	insertAccountStmt           = "INSERT INTO ACCOUNTS (USER_ID, BALANCE, CREATION_DATE, MODIFICATION_DATE)VALUES (?,?,?,?)"
	selectAccountByUsernameStmt = "SELECT A.ID, A.USER_ID, A.BALANCE, A.CREATION_DATE, A.MODIFICATION_DATE FROM ACCOUNTS A JOIN USERS U ON U.ID = A.USER_ID WHERE U.USERNAME = ?"
	updateAccountBalanceStmt    = "UPDATE ACCOUNTS SET BALANCE = ? WHERE ID = ?"
)

var (
	AccountNotFoundError = api_error.NewApiError(http.StatusNotFound, "account not found")
)

type Account interface {
	SaveAccount(ctx context.Context, account AccountEntity) error
	GetAccountByUsername(ctx context.Context, username string) (*AccountEntity, error)
	UpdateAccountBalance(ctx context.Context, account AccountEntity) error
}

type AccountDependencies struct {
	Database *sql.DB
}

type AccountImpl struct {
	database *sql.DB
}

func NewAccountImpl(dependencies AccountDependencies) AccountImpl {
	return AccountImpl{
		database: dependencies.Database,
	}
}

type AccountEntity struct {
	Id               int
	UserId           int
	Balance          float64
	CreationDate     string
	ModificationDate string
}

func (r AccountImpl) SaveAccount(ctx context.Context, account AccountEntity) error {
	timeNow := time.Now()

	_, err := r.database.ExecContext(ctx, insertAccountStmt, account.UserId, initialAccountBalance, timeNow, timeNow)
	if err != nil {
		return err
	}

	return nil
}

func (r AccountImpl) GetAccountByUsername(ctx context.Context, username string) (*AccountEntity, error) {
	row := r.database.QueryRowContext(ctx, selectAccountByUsernameStmt, username)

	var account AccountEntity
	err := row.Scan(&account.Id, &account.UserId, &account.Balance, &account.CreationDate, &account.ModificationDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, AccountNotFoundError
		}
		return nil, err
	}

	return &account, nil
}

func (r AccountImpl) UpdateAccountBalance(ctx context.Context, account AccountEntity) error {
	_, err := r.database.ExecContext(ctx, updateAccountBalanceStmt, account.Balance, account.Id)

	if err != nil {
		return err
	}

	return nil
}