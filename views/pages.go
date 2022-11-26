package views

import (
	"html/template"
	"net/http"

	"github.com/cyberphor/ctfconsole/controllers"
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
	filePath := http.Dir("views/")
	fileServer := http.FileServer(filePath)
	http.Handle("/static/", fileServer)
	http.HandleFunc("/admin/index.html", AdminLoginPage)
	http.HandleFunc("/index.html", LoginPage)
	http.HandleFunc("/register.html", RegisterPage)
	http.HandleFunc("/login", controllers.Login)
	//http.HandleFunc("/register", Register)
	http.HandleFunc("/admin/login", controllers.AdminLogin)
	http.HandleFunc("/challenges.html", controllers.VerifyToken(ScoreboardPage))
	http.HandleFunc("/scoreboard.html", controllers.VerifyToken(ScoreboardPage))
	http.HandleFunc("/admin/create-ctf.html", controllers.VerifyToken(CreateCtfPage))
	http.HandleFunc("/", LoginPage)
	http.ListenAndServe(":8000", nil)
}
