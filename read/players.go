package read

import (
	"database/sql"

	"github.com/cyberphor/ctfconsole/models"
)

func Players() []models.PlayerData {
	database, err := sql.Open("sqlite3", "ctfconsole.db")
	if err != nil {
		panic(err)
	}
	statement := `SELECT username,password FROM players;`
	rows, err := database.Query(statement)
	if err != nil {
		panic(err)
	}
	var Players []models.PlayerData
	var username string
	var password string
	for rows.Next() {
		rows.Scan(&username, &password)
		Player := models.PlayerData{
			Username: username,
			Password: password,
		}
		Players = append(Players, Player)
	}
	database.Close()
	return Players
}
