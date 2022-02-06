package controllers

import (
	"net/http"

	"github.com/cyberphor/ctfconsole/models"
)

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
