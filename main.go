package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"html/template"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	_ "github.com/mattn/go-sqlite3"
)

var templates *template.Template

var key = []byte("The true crimefighter always carries everything he needs in his utility belt, Robin.")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type User struct {
	Username string
	Password string
	Team     string
	Role     string
}

func init() {
	templates = template.Must(template.ParseGlob("templates/*"))
}

func AdminLoginPage(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "admin-login.gohtml", nil)
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

func Console() {
	filePath := http.Dir("/")
	fileServer := http.FileServer(filePath)
	http.Handle("/static/", fileServer)
	http.HandleFunc("/admin/index.html", AdminLoginPage)
	http.HandleFunc("/admin/login", AdminLogin)
	http.HandleFunc("/admin/create-ctf.html", VerifyToken(CreateCtfPage))
	http.HandleFunc("/index.html", LoginPage)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/register.html", RegisterPage)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/challenges.html", VerifyToken(ScoreboardPage))
	http.HandleFunc("/scoreboard.html", VerifyToken(ScoreboardPage))
	http.HandleFunc("/", LoginPage)
	http.ListenAndServe(":8000", nil)
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

func CreateCtfPage(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create-ctf.gohtml", nil)
}

func CreateTableForAdmins() {
	database, err := sql.Open("sqlite3", "accounts.db")
	if err != nil {
		panic(err)
	}
	statement := `CREATE TABLE IF NOT EXISTS admins (
		id INTEGER PRIMARY KEY,
		username  TEXT,
		password  TEXT,
		role      TEXT,
		UNIQUE(username)
		);`
	query, err := database.Prepare(statement)
	if err != nil {
		panic(err)
	}
	query.Exec()
	database.Close()
}

func CreateTableForUsers() {
	database, err := sql.Open("sqlite3", "accounts.db")
	if err != nil {
		panic(err)
	}
	statement := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		username  TEXT,
		password  TEXT,
		team      TEXT,
		role      TEXT,
		UNIQUE(username)
		);`
	query, err := database.Prepare(statement)
	if err != nil {
		panic(err)
	}
	query.Exec()
	database.Close()
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

func GetPasswordHash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
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

func LoginPage(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "user-login.gohtml", nil)
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

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "register.gohtml", nil)
}

func ScoreboardPage(w http.ResponseWriter, r *http.Request) {
	type data struct {
		Authenticated bool
		Users         []User
	}
	templateData := data{
		Authenticated: true,
		Users:         GetUsers(),
	}
	templates.ExecuteTemplate(w, "scoreboard.gohtml", templateData)
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

func main() {
	CreateTableForUsers()
	CreateTableForAdmins()
	CreateAdmin("admin", "password", "admin")
	Console()
}
