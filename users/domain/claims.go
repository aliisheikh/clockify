package domain

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
