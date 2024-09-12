package clients

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

const (
	dbPath = "db/db"
)

var (
	creationSentences = map[string]string{
		"users": `CREATE TABLE IF NOT EXISTS USERS(
						ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
						USERNAME TEXT UNIQUE,
						PASSWORD TEXT,
						STATUS TEXT,
						LOGIN_ATTEMPTS INTEGER,
						MODIFICATION_DATE TEXT,
						CREATION_DATE TEXT
					)`,
		"accounts": `CREATE TABLE IF NOT EXISTS ACCOUNTS(
						ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
						USER_ID INTEGER,
						BALANCE REAL,
						CURRENCY TEXT,
						CREATION_DATE TEXT,
						MODIFICATION_DATE TEXT,
						FOREIGN KEY (USER_ID) REFERENCES USERS(ID)
					)`,
		"historic": `CREATE TABLE IF NOT EXISTS TRANSACTIONS(
						ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
						USER_ID INTEGER,
						ACCOUNT_ID INTEGER,
						AMOUNT REAL,
						PARTIAL_BALANCE REAL,
						TYPE TEXT,
						DATE TEXT,
						FOREIGN KEY (USER_ID) REFERENCES USERS(ID),
						FOREIGN KEY (ACCOUNT_ID) REFERENCES ACCOUNTS(ID)
					)`,
	}
)

func NewSQLite() (*sql.DB, error) {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		dir := filepath.Dir(dbPath)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}

	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	for _, sentence := range creationSentences {
		query, err := database.Prepare(sentence)
		if err != nil {
			database.Close()
			return nil, err
		}

		_, err = query.Exec()
		if err != nil {
			database.Close()
			return nil, err
		}
	}

	return database, nil
}