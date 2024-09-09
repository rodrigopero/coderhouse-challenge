package clients

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var (
	dbFilename        = "db.db"
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

func InitializeDbSqlite(path string) (*sql.DB, error) {
	if _, err := os.Stat(dbFilename); os.IsNotExist(err) {
		_, err := os.Create(dbFilename)

		if err != nil {
			return nil, err
		}

	}

	database, err := sql.Open("sqlite3", dbFilename)
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