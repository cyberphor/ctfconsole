package controllers

import (
	"html/template"
	"net/http"

	"github.com/cyberphor/ctfconsole/models"
)

func Scoreboard(w http.ResponseWriter, r *http.Request) {
	pageTemplate, err := template.ParseFiles("./views/scoreboard.gohtml")
	if err != nil {
		panic(err)
	}
	err = pageTemplate.Execute(w, models.GetPlayers())
	if err != nil {
		panic(err)
	}
}
