package models

import "database/sql"

func AuthenticateAdmin(username string, password string) bool {
	database, err := sql.Open("sqlite3", "store/store.db")
	if err != nil {
		panic(err)
	}
	statement := `SELECT id FROM admins WHERE (username,password) = (?,?);`
	row := database.QueryRow(statement, username, GetPasswordHash(password))
	var id string
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		return false
	case nil:
		return true
	default:
		panic(err)
	}
}

func CreateAdmin(username string, password string, role string) {
	database, err := sql.Open("sqlite3", "store/store.db")
	if err != nil {
		panic(err)
	}
	statement := `INSERT OR IGNORE INTO admins (username, password, role) VALUES (?,?,?);`
	query, err := database.Prepare(statement)
	if err != nil {
		panic(err)
	}
	query.Exec(username, GetPasswordHash(password), role)
	database.Close()
}
