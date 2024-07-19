package application

import (
	"clockify/users/domain"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func (u *UserServiceImp) GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &domain.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(u.jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
