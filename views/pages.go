package views

import (
	"html/template"
	"net/http"

	"github.com/cyberphor/ctfconsole/models"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("views/templates/*"))
}

func AdminLoginPage(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "admin-login.gohtml", nil)
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "register.gohtml", nil)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "user-login.gohtml", nil)
}

func CreateCtfPage(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create-ctf.gohtml", nil)
}

func ScoreboardPage(w http.ResponseWriter, r *http.Request) {
	type data struct {
		Authenticated bool
		Users         []models.User
	}
	templateData := data{
		Authenticated: true,
		Users:         models.GetUsers(),
	}
	templates.ExecuteTemplate(w, "scoreboard.gohtml", templateData)
}

func Console() {
	filePath := http.Dir("/")
	fileServer := http.FileServer(filePath)
	http.Handle("/static/", fileServer)
	http.HandleFunc("/admin/index.html", AdminLoginPage)
	//http.HandleFunc("/admin/login", AdminLogin)
	//http.HandleFunc("/admin/create-ctf.html", VerifyToken(CreateCtfPage))
	http.HandleFunc("/index.html", LoginPage)
	//http.HandleFunc("/login", Login)
	http.HandleFunc("/register.html", RegisterPage)
	//http.HandleFunc("/register", Register)
	//http.HandleFunc("/challenges.html", VerifyToken(ScoreboardPage))
	//http.HandleFunc("/scoreboard.html", VerifyToken(ScoreboardPage))
	http.HandleFunc("/", LoginPage)
	http.ListenAndServe(":8000", nil)
}
