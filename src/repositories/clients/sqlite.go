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
		"users": `CREATE TABLE IF NOT EXISTS users(
						id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
						username TEXT,
						password TEXT,
						status TEXT,
						creation_date TEXT,
						modification_date TEXT
					)`,
		"accounts": `CREATE TABLE IF NOT EXISTS accounts(
						id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
						user_id INTEGER,
						currency TEXT,
						balance REAL,
						status TEXT,
						creation_date TEXT,
						modification_date TEXT,
						FOREIGN KEY (user_id) REFERENCES users(id)
					)`,
		"historic": `CREATE TABLE IF NOT EXISTS historic(
						id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
						user_id INTEGER,
						account_id INTEGER,
						amount REAL,
						status TEXT,
						creation_date TEXT,
						modification_date TEXT,
						FOREIGN KEY (user_id) REFERENCES users(id),
						FOREIGN KEY (account_id) REFERENCES accounts(id)
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
