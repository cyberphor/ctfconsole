package controllers

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("views/templates/*"))
}

func Console() {
	filePath := http.Dir("views/")
	fileServer := http.FileServer(filePath)
	http.Handle("/static/", fileServer)
	http.HandleFunc("/", LoginPage)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/register.html", RegisterPage)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/challenges.html", VerifyToken(ScoreboardPage))
	http.HandleFunc("/scoreboard.html", VerifyToken(ScoreboardPage))
	http.HandleFunc("/admin/create-ctf.html", VerifyToken(CreateCtfPage))
	http.ListenAndServe(":8000", nil)
}
