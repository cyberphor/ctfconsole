package create

import "database/sql"

func Player(username string, password string, team string) {
	database, err := sql.Open("sqlite3", "ctfconsole.db")
	if err != nil {
		panic(err)
	}
	statement := `INSERT OR IGNORE INTO players (username, password, team) VALUES (?,?,?);`
	query, err := database.Prepare(statement)
	if err != nil {
		panic(err)
	}
	query.Exec(username, HashPassword(password), team)
	database.Close()
}
