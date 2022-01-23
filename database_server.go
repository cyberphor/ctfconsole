package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"

	_ "github.com/mattn/go-sqlite3"
)

func HashPassword(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func CreateAdminsTable() {
	database, err := sql.Open("sqlite3", "./ctfconsole.db")
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

func CreateAdmin(username string, password string) {
	database, err := sql.Open("sqlite3", "./ctfconsole.db")
	if err != nil {
		panic(err)
	}
	statement := `INSERT OR IGNORE INTO admins (username, password) VALUES (?,?);`
	query, err := database.Prepare(statement)
	if err != nil {
		panic(err)
	}
	query.Exec(username, HashPassword(password))
	database.Close()
}

func ReadAdminsTable() {
	database, err := sql.Open("sqlite3", "./ctfconsole.db")
	if err != nil {
		panic(err)
	}
	statement := `SELECT username,password FROM admins;`
	rows, err := database.Query(statement)
	if err != nil {
		panic(err)
	} else {
		database.Close()
		return rows
	}
	/*
		var username string
		var password string
		for rows.Next() {
			rows.Scan(&username, &password)
	*/
}

func CreatePlayersTable() {
	database, err := sql.Open("sqlite3", "./ctfconsole.db")
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

func CreatePlayer(username string, password string) {
	database, err := sql.Open("sqlite3", "./ctfconsole.db")
	if err != nil {
		panic(err)
	}
	statement := `INSERT OR IGNORE INTO players (username, password) VALUES (?,?);`
	query, err := database.Prepare(statement)
	if err != nil {
		panic(err)
	}
	query.Exec(username, HashPassword(password))
	database.Close()
}

func ReadPlayersTable() {
	database, err := sql.Open("sqlite3", "./ctfconsole.db")
	if err != nil {
		panic(err)
	}
	statement := `SELECT username,password FROM players;`
	rows, err := database.Query(statement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	type Player struct {
		Username string
		Password string
	}
	var players []Player
	var username string
	var password string
	for rows.Next() {
		var row Player
		rows.Scan(&row.username, &row.password)
		players = append(players, row)
	}
	return players
}
