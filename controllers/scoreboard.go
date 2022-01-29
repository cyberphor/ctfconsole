package controllers

import (
	"html/template"
	"net/http"

	"github.com/cyberphor/ctfconsole/models"
)

func showScoreboard(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("./templates/scoreboard.html")
	if err != nil {
		panic(err)
	}
	players := models.GetPlayers()
	page.Execute(w, players)
}
