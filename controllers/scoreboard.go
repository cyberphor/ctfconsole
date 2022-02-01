package controllers

import (
	"net/http"

	"github.com/cyberphor/ctfconsole/models"
)

func Scoreboard(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "scoreboard.gohtml", models.GetPlayers())
}
