package main

import (
	"database/sql"
	"net/http"
)

type User struct {
	Username string
	Password string
	Team     string
	Role     string
}

func CreateAdmin(username string, password string, role string) {
	database, err := sql.Open("sqlite3", "accounts.db")
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

func CreateUser(username string, password string, team string, role string) {
	database, err := sql.Open("sqlite3", "accounts.db")
	if err != nil {
		panic(err)
	}
	statement := `INSERT OR IGNORE INTO users (username, password, team, role) VALUES (?,?,?,?);`
	query, err := database.Prepare(statement)
	if err != nil {
		panic(err)
	}
	query.Exec(username, GetPasswordHash(password), team, role)
	database.Close()
}

func GetUsers() []User {
	database, err := sql.Open("sqlite3", "accounts.db")
	if err != nil {
		panic(err)
	}
	statement := `SELECT username, team, role FROM users;`
	rows, err := database.Query(statement)
	if err != nil {
		panic(err)
	}
	var users []User
	var username string
	var team string
	var role string
	for rows.Next() {
		rows.Scan(&username, &team, &role)
		user := User{
			Username: username,
			Team:     team,
			Role:     role,
		}
		users = append(users, user)
	}
	database.Close()
	return users
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		CreateUser(r.FormValue("username"), r.FormValue("password"), r.FormValue("team"), "user")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}
