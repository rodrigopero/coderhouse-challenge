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

	insertAccountStmt                      = "INSERT INTO ACCOUNTS (USER_ID, BALANCE, CREATION_DATE, MODIFICATION_DATE, CURRENCY)VALUES (?,?,?,?,?)"
	selectAccountByUsernameAndCurrencyStmt = "SELECT A.ID, A.USER_ID, A.BALANCE, A.CREATION_DATE, A.MODIFICATION_DATE, A.CURRENCY FROM ACCOUNTS A JOIN USERS U ON U.ID = A.USER_ID WHERE U.USERNAME = ? AND A.CURRENCY = ?"
	selectAccountByUsernameStmt            = "SELECT A.ID, A.USER_ID, A.BALANCE, A.CREATION_DATE, A.MODIFICATION_DATE, A.CURRENCY FROM ACCOUNTS A JOIN USERS U ON U.ID = A.USER_ID WHERE U.USERNAME = ?"
	updateAccountBalanceStmt               = "UPDATE ACCOUNTS SET BALANCE = ? WHERE ID = ?  AND CURRENCY = ?"
)

var (
	AccountNotFoundError = api_error.NewApiError(http.StatusNotFound, "Account not found")
)

type Account interface {
	SaveAccount(ctx context.Context, account AccountEntity) error
	GetAccountByUsernameAndCurrency(ctx context.Context, username string, currency string) (*AccountEntity, error)
	GetAccountsByUsername(ctx context.Context, username string) ([]AccountEntity, error)
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
	Currency         string
	CreationDate     string
	ModificationDate string
}

func (r AccountImpl) SaveAccount(ctx context.Context, account AccountEntity) error {
	timeNow := time.Now()

	_, err := r.database.ExecContext(ctx, insertAccountStmt, account.UserId, initialAccountBalance, timeNow, timeNow, account.Currency)
	if err != nil {
		return UnexpectedError
	}

	return nil
}

func (r AccountImpl) GetAccountByUsernameAndCurrency(ctx context.Context, username string, currency string) (*AccountEntity, error) {
	row := r.database.QueryRowContext(ctx, selectAccountByUsernameAndCurrencyStmt, username, currency)

	var account AccountEntity
	err := row.Scan(&account.Id, &account.UserId, &account.Balance, &account.CreationDate, &account.ModificationDate, &account.Currency)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, AccountNotFoundError
		}
		return nil, UnexpectedError
	}

	return &account, nil
}

func (r AccountImpl) GetAccountsByUsername(ctx context.Context, username string) ([]AccountEntity, error) {
	rows, err := r.database.QueryContext(ctx, selectAccountByUsernameStmt, username)
	if err != nil {
		return nil, UnexpectedError
	}
	defer rows.Close()

	var accountList []AccountEntity

	for rows.Next() {
		var account AccountEntity
		err = rows.Scan(&account.Id, &account.UserId, &account.Balance, &account.CreationDate, &account.ModificationDate, &account.Currency)
		if err != nil {
			return nil, UnexpectedError
		}

		accountList = append(accountList, account)
	}

	if err = rows.Err(); err != nil {
		return nil, UnexpectedError
	}

	return accountList, nil
}

func (r AccountImpl) UpdateAccountBalance(ctx context.Context, account AccountEntity) error {
	_, err := r.database.ExecContext(ctx, updateAccountBalanceStmt, account.Balance, account.Id, account.Currency)

	if err != nil {
		return UnexpectedError
	}

	return nil
}
