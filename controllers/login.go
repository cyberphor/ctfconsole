package controllers

import (
	"fmt"
	"net/http"

	"github.com/cyberphor/ctfconsole/models"
)

func Auth(roles []string, HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			// http.StatusBadRequest
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		token, claims, err := ParseTokenString(cookie.Value)
		if err != nil {
			// http.StatusBadRequest
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		if token.Valid {
			fmt.Println(roles, claims.Role)
			for _, role := range roles {
				if claims.Role == role {
					HandlerFunc.ServeHTTP(w, r)
				} else {
					// http.StatusUnauthorized
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
			}
		} else {
			// http.StatusUnauthorized
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if models.GetPlayer(username, password) != "failed" {
			tokenString, expirationTime, err := CreateToken(username, "user")
			if err != nil {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			http.SetCookie(w,
				&http.Cookie{
					Name:    "token",
					Value:   tokenString,
					Expires: expirationTime,
				},
			)
			http.Redirect(w, r, "/scoreboard.html", http.StatusSeeOther)
			return
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login.gohtml", nil)
}
