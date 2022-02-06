package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTableForUsers() {
	database, err := sql.Open("sqlite3", "store/store.db")
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

func CreateTableForScoreboard() {
	/*
		database, err := sql.Open("sqlite3", "store/store.db")
		if err != nil {
			panic(err)
		}
		statement := ``
		query, err := database.Prepare(statement)
		if err != nil {
			panic(err)
		}
		query.Exec()
		database.Close()
	*/
}
