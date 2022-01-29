package controllers

import (
	"net/http"

	"github.com/cyberphor/ctfconsole/models"
)

func submitLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if models.GetPlayer(username, password) != "failed" {
			cookie, err := r.Cookie("session")
			if err != nil {
				cookie = &http.Cookie{
					Name:  "session",
					Value: "123456789",
				}
				http.SetCookie(w, cookie)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		} else {
			http.Redirect(w, r, "/login.html", http.StatusSeeOther)
			return
		}
	}
}
