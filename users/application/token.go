package application

import (
	"time"

	"clockify/users/domain"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("yoursecretkey")

func (u *UserServiceImp) GenerateToken(userID int, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &domain.Claims{
		ID:       userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
