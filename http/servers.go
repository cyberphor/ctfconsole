package ctfconsole

import (
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var key = []byte("The true crimefighter always carries everything he needs in his utility belt, Robin.")

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

func main() {
	CreateTableForUsers()
	CreateTableForAdmins()
	CreateAdmin("admin", "password", "admin")
	Console()
}
