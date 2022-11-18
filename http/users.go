package ctfconsole

import "github.com/golang-jwt/jwt"

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type User struct {
	Username string
	Password string
	Team     string
	Role     string
}
