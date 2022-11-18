package ctfconsole

import (
	"database/sql"
	"net/http"
)

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

func AuthenticateAdmin(username string, password string) bool {
	database, err := sql.Open("sqlite3", "accounts.db")
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

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if AuthenticateAdmin(username, password) {
			tokenString, expirationTime, err := CreateToken(username, "admins")
			if err != nil {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			http.SetCookie(w,
				&http.Cookie{
					Name:    "token",
					Value:   tokenString,
					Expires: expirationTime,
				},
			)
			http.Redirect(w, r, "/scoreboard.html", http.StatusSeeOther)
			return
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
}
