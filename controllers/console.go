package controllers

import (
	"net/http"
)

func Console() {
	filePath := http.Dir("./views/")
	fileServer := http.FileServer(filePath)
	http.Handle("/static/", fileServer)
	http.HandleFunc("/", LoginPage)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/scoreboard.html", Scoreboard)
	http.ListenAndServe(":8000", nil)
}
