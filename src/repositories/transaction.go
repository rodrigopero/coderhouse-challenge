package repositories

import (
	"context"
	"database/sql"
	"time"
)

const (
	insertTransactionStmt        = `INSERT INTO TRANSACTIONS (USER_ID, ACCOUNT_ID, AMOUNT, PARTIAL_BALANCE, TYPE, "DATE") VALUES (?,?,?,?,?,?)`
	getTransactionsWithLimitStmt = `SELECT T.ID, T.USER_ID, T.ACCOUNT_ID, T.AMOUNT, T.PARTIAL_BALANCE, T.TYPE, T."DATE" FROM TRANSACTIONS T JOIN MAIN.USERS U ON U.ID = T.USER_ID WHERE U.USERNAME = ? AND T.ACCOUNT_ID = ? ORDER BY DATE DESC LIMIT ?`
)

type Transaction interface {
	SaveTransaction(ctx context.Context, transaction TransactionEntity) error
	GetTransactionsWithLimit(ctx context.Context, username string, accountID int, limit int) ([]TransactionEntity, error)
}

type TransactionImpl struct {
	database *sql.DB
}

type TransactionDependencies struct {
	Database *sql.DB
}

func NewTransactionImpl(dependencies TransactionDependencies) TransactionImpl {
	return TransactionImpl{
		database: dependencies.Database,
	}
}

type TransactionEntity struct {
	Id             int
	UserId         int
	AccountId      int
	Amount         float64
	PartialBalance float64
	Type           string
	Date           string
}

func (r TransactionImpl) SaveTransaction(ctx context.Context, transaction TransactionEntity) error {
	timeNow := time.Now()

	_, err := r.database.ExecContext(ctx, insertTransactionStmt, transaction.UserId, transaction.AccountId,
		transaction.Amount, transaction.PartialBalance, transaction.Type, timeNow, timeNow)
	if err != nil {
		return err
	}

	return nil
}

func (r TransactionImpl) GetTransactionsWithLimit(ctx context.Context, username string, accountID int, limit int) ([]TransactionEntity, error) {
	rows, err := r.database.QueryContext(ctx, getTransactionsWithLimitStmt, username, accountID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactionList []TransactionEntity

	for rows.Next() {
		var transaction TransactionEntity
		err := rows.Scan(&transaction.Id, &transaction.UserId, &transaction.AccountId, &transaction.Amount, &transaction.PartialBalance, &transaction.Type, &transaction.Date)
		if err != nil {
			return nil, err
		}

		transactionList = append(transactionList, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactionList, nil
}
