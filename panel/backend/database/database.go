package database

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init() error {

	var err error

	DB, err = sql.Open("sqlite", "data.db")

	if err != nil {
		return err
	}

	query := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		uuid TEXT,
		volume_limit INTEGER,
		volume_used INTEGER,
		expire_date TEXT,
		max_connections INTEGER,
		status TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = DB.Exec(query)

	return err
}