package controllers

import (
	"net/http"

	"github.com/cyberphor/ctfconsole/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		models.CreateUser(r.FormValue("username"), r.FormValue("password"), r.FormValue("team"), "user")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "register.gohtml", nil)
}
