package ctfconsole

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("templates/*"))
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
		Users         []User
	}
	templateData := data{
		Authenticated: true,
		Users:         GetUsers(),
	}
	templates.ExecuteTemplate(w, "scoreboard.gohtml", templateData)
}
