package models

import "database/sql"

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

func GetPlayer(username string, password string) string {
	database, err := sql.Open("sqlite3", "store/store.db")
	if err != nil {
		panic(err)
	}
	statement := `SELECT username FROM players WHERE (username,password) = (?,?);`
	row := database.QueryRow(statement, username, GetPasswordHash(password))
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
