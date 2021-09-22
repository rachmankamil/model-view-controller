package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rachmankamil/kampus-merdeka-b/config"
)

func CreateToken(userID int) (token string, err error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = userID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(config.JWT_EXP)).Unix()

	initToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = initToken.SignedString(config.JWT_SECRET)
	return
}
