package controllers

import (
	"net/http"
	"time"

	"github.com/cyberphor/ctfconsole/models"
	"github.com/golang-jwt/jwt"
)

var key = []byte("The true crimefighter always carries everything he needs in his utility belt, Robin.")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Auth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// http.StatusBadRequest
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			// http.StatusBadRequest
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		tokenString := cookie.Value
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims,
			func(t *jwt.Token) (interface{}, error) {
				return key, nil
			})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				// http.StatusUnauthorized
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			// http.StatusBadRequest
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		if !token.Valid {
			// http.StatusUnauthorized
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		HandlerFunc.ServeHTTP(w, r)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		// check credentials
		credentials := &Credentials{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		if models.GetPlayer(credentials.Username, credentials.Password) != "failed" {
			// create a JWT
			expirationTime := time.Now().Add(time.Minute * 5)
			claims := &Claims{
				Username: credentials.Username,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(key)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			// return the JWT as a cookie
			http.SetCookie(w,
				&http.Cookie{
					Name:    "token",
					Value:   tokenString,
					Expires: expirationTime,
				},
			)
			http.Redirect(w, r, "/scoreboard.html", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login.gohtml", nil)
}
