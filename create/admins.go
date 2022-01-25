package create

import "database/sql"

func Admin(username string, password string) {
	database, err := sql.Open("sqlite3", "ctfconsole.db")
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
