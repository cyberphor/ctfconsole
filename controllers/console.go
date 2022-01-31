package controllers

import (
	"net/http"
)

func Console() {
	filePath := http.Dir("./views/")
	fileServer := http.FileServer(filePath)
	http.Handle("/static/", fileServer)
	http.HandleFunc("/", LoginPage)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/scoreboard.html", showScoreboard)
	http.HandleFunc("/submitRegistration", submitRegistration)
	http.ListenAndServe(":8000", nil)
}
