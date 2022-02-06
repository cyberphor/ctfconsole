package controllers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var key = []byte("The true crimefighter always carries everything he needs in his utility belt, Robin.")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func CreateToken(username string, role string) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	return tokenString, expirationTime, err
}

func ParseTokenString(tokenString string) (*jwt.Token, *Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	claims := token.Claims.(*Claims)
	return token, claims, err
}

func VerifyToken(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			// http.StatusBadRequest
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		token, _, err := ParseTokenString(cookie.Value)
		if err != nil {
			// http.StatusBadRequest
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		if token.Valid {
			HandlerFunc.ServeHTTP(w, r)
		} else {
			// http.StatusUnauthorized
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
}
