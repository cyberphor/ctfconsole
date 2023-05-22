package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTableForAdmins() {
	database, err := sql.Open("sqlite3", "accounts.db")
	if err != nil {
		panic(err)
	}
	statement := `CREATE TABLE IF NOT EXISTS admins (
		id INTEGER PRIMARY KEY,
		username  TEXT,
		password  TEXT,
		role      TEXT,
		UNIQUE(username)
		);`
	query, err := database.Prepare(statement)
	if err != nil {
		panic(err)
	}
	query.Exec()
	database.Close()
}

func CreateTableForUsers() {
	database, err := sql.Open("sqlite3", "accounts.db")
	if err != nil {
		panic(err)
	}
	statement := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		username  TEXT,
		password  TEXT,
		team      TEXT,
		role      TEXT,
		UNIQUE(username)
		);`
	query, err := database.Prepare(statement)
	if err != nil {
		panic(err)
	}
	query.Exec()
	database.Close()
}
