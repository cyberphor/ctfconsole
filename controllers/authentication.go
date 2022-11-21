package controllers

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

var key = []byte("The true crimefighter always carries everything he needs in his utility belt, Robin.")

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

func AuthenticateUser(username string, password string) bool {
	database, err := sql.Open("sqlite3", "accounts.db")
	if err != nil {
		panic(err)
	}
	statement := `SELECT id FROM users WHERE (username,password) = (?,?);`
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

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if AuthenticateUser(username, password) {
			tokenString, expirationTime, err := CreateToken(username, "user")
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

func GetPasswordHash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func CreateToken(username string, role string) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	return tokenString, expirationTime, err
}

func ParseTokenString(tokenString string) (*jwt.Token, *Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	claims := token.Claims.(*Claims)
	return token, claims, err
}

func VerifyToken(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		token, _, err := ParseTokenString(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		if token.Valid {
			HandlerFunc.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
}
