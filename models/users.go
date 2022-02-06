package models

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
)

type User struct {
	Username string
	Password string
	Team     string
	Role     string
}

func AuthenticateUser(username string, password string) bool {
	database, err := sql.Open("sqlite3", "store/store.db")
	if err != nil {
		panic(err)
	}
	statement := `SELECT id FROM users WHERE (username,password) = (?,?);`
	row := database.QueryRow(statement, username, GetPasswordHash(password))
	var id string
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		fmt.Println("false")
		return false
	case nil:
		return true
	default:
		panic(err)
	}
}

func CreateUser(username string, password string, team string, role string) {
	database, err := sql.Open("sqlite3", "store/store.db")
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

func GetPasswordHash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetUsers() []User {
	database, err := sql.Open("sqlite3", "store/store.db")
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
