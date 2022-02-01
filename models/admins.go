package models

import "database/sql"

func CreateAdmin(username string, password string) {
	database, err := sql.Open("sqlite3", "store/store.db")
	if err != nil {
		panic(err)
	}
	statement := `INSERT OR IGNORE INTO admins (username, password) VALUES (?,?);`
	query, err := database.Prepare(statement)
	if err != nil {
		panic(err)
	}
	query.Exec(username, GetPasswordHash(password))
	database.Close()
}
