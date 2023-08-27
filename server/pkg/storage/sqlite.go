package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CreateSQLiteDatabase(name string) {
	var db *sql.DB
	var err error
	var queries []string
	var query string
	var statement *sql.Stmt

	db, err = sql.Open("sqlite3", name)
	if err != nil {
		panic(err)
	}

	queries = []string{
		`CREATE TABLE IF NOT EXISTS admins (
			id INTEGER PRIMARY KEY,
			name  TEXT,
			password  TEXT,
			UNIQUE(name)
		);`,
		`CREATE TABLE IF NOT EXISTS challenges (
			id INTEGER PRIMARY KEY,
			name      TEXT,
			points    TEXT,
			solution  TEXT,
			UNIQUE(name)
		);`,
		`CREATE TABLE IF NOT EXISTS players (
			id INTEGER PRIMARY KEY,
			name  TEXT,
			password  TEXT,
			UNIQUE(name)
		);`,
		`CREATE TABLE IF NOT EXISTS scoreboards (
			id INTEGER PRIMARY KEY,
			name      TEXT,
			teams     TEXT,
			UNIQUE(name)
		);`,
		`CREATE TABLE IF NOT EXISTS team (
			id INTEGER PRIMARY KEY,
			name      TEXT,
			players   TEXT,
			UNIQUE(name)
		);`,
	}

	for _, query = range queries {
		statement, err = db.Prepare(query)
		if err != nil {
			panic(err)
		}
		statement.Exec()
	}
	db.Close()
}
