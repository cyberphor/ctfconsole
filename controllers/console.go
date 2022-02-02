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
	http.HandleFunc("/scoreboard.html", Scoreboard)
	http.HandleFunc("/challenges.html", CheckToken)
	http.ListenAndServe(":8000", nil)
}
