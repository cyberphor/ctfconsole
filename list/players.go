package list

import "database/sql"

type PlayerData struct {
	Username string
	Password string
}

func ListPlayers() []PlayerData {
	database, err := sql.Open("sqlite3", "./ctfconsole.db")
	if err != nil {
		panic(err)
	}
	statement := `SELECT username,password FROM players;`
	rows, err := database.Query(statement)
	if err != nil {
		panic(err)
	}
	var Players []PlayerData
	var username string
	var password string
	for rows.Next() {
		rows.Scan(&username, &password)
		Player := PlayerData{
			Username: username,
			Password: password,
		}
		Players = append(Players, Player)
	}
	database.Close()
	return Players
}
