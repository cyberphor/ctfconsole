package controllers

import (
	"net/http"
)

func Console() {
	filePath := http.Dir("./templates/")
	fileServer := http.FileServer(filePath)
	http.Handle("/", fileServer)
	http.HandleFunc("/scoreboard.html", showScoreboard)
	http.HandleFunc("/submitRegistration", submitRegistration)
	http.HandleFunc("/submitLogin", submitLogin)
	http.ListenAndServe(":8000", nil)
}
