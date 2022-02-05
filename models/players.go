package models

import (
	"database/sql"
	"fmt"
)

type PlayerData struct {
	Username string
	Password string
	Team     string
}

func CreatePlayer(username string, password string, team string) {
	database, err := sql.Open("sqlite3", "store/store.db")
	if err != nil {
		panic(err)
	}
	statement := `INSERT OR IGNORE INTO players (username, password, team) VALUES (?,?,?);`
	query, err := database.Prepare(statement)
	if err != nil {
		panic(err)
	}
	query.Exec(username, GetPasswordHash(password), team)
	database.Close()
}

func GetPlayers() []PlayerData {
	database, err := sql.Open("sqlite3", "store/store.db")
	if err != nil {
		panic(err)
	}
	statement := `SELECT username,password,team FROM players;`
	rows, err := database.Query(statement)
	if err != nil {
		panic(err)
	}
	var Players []PlayerData
	var username string
	var password string
	var team string
	for rows.Next() {
		rows.Scan(&username, &password, &team)
		Player := PlayerData{
			Username: username,
			Password: password,
			Team:     team,
		}
		Players = append(Players, Player)
	}
	database.Close()
	return Players
}

func AuthenticatePlayer(username string, password string) bool {
	database, err := sql.Open("sqlite3", "store/store.db")
	if err != nil {
		panic(err)
	}
	statement := `SELECT id FROM players WHERE (username,password) = (?,?);`
	fmt.Println(username)
	row := database.QueryRow(statement, username, GetPasswordHash(password))
	var id string
	fmt.Print(row.Scan(&id))
	err = row.Scan()
	if err != nil {
		return false
	} else {
		return true
	}
}
