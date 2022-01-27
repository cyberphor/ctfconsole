package read

import (
	"database/sql"

	"github.com/cyberphor/ctfconsole/create"
	"github.com/cyberphor/ctfconsole/models"
)

func Players() []models.PlayerData {
	database, err := sql.Open("sqlite3", "ctfconsole.db")
	if err != nil {
		panic(err)
	}
	statement := `SELECT username,password,team FROM players;`
	rows, err := database.Query(statement)
	if err != nil {
		panic(err)
	}
	var Players []models.PlayerData
	var username string
	var password string
	var team string
	for rows.Next() {
		rows.Scan(&username, &password, &team)
		Player := models.PlayerData{
			Username: username,
			Password: password,
			Team:     team,
		}
		Players = append(Players, Player)
	}
	database.Close()
	return Players
}

func Player(username string, password string) string {
	database, err := sql.Open("sqlite3", "ctfconsole.db")
	if err != nil {
		panic(err)
	}
	statement := `SELECT username FROM players WHERE (username,password) = (?,?);`
	row := database.QueryRow(statement, username, create.HashPassword(password))
	var PlayerUsername string
	switch err := row.Scan(&PlayerUsername); err {
	case nil:
		return PlayerUsername
	case sql.ErrNoRows:
		return "failed"
	default:
		panic(err)
	}
}
