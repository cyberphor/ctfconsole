package create

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func TableForAdmins() {
	database, err := sql.Open("sqlite3", "ctfconsole.db")
	if err != nil {
		panic(err)
	}
	statement := `CREATE TABLE IF NOT EXISTS admins (
		id INTEGER PRIMARY KEY,
		username  TEXT,
		password  TEXT,
		UNIQUE(username)
		);`
	query, err := database.Prepare(statement)
	if err != nil {
		panic(err)
	}
	query.Exec()
	database.Close()
}

func TableForPlayers() {
	database, err := sql.Open("sqlite3", "ctfconsole.db")
	if err != nil {
		panic(err)
	}
	statement := `CREATE TABLE IF NOT EXISTS players (
		id INTEGER PRIMARY KEY,
		username  TEXT,
		password  TEXT,
		UNIQUE(username)
		);`
	query, err := database.Prepare(statement)
	if err != nil {
		panic(err)
	}
	query.Exec()
	database.Close()
}
