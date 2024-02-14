package utils_jwt

import (
	log "github.com/sirupsen/logrus"

	"config"

	"errors"

	"github.com/golang-jwt/jwt"
	//"time"
)

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	//claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = username

	tokenString, err := token.SignedString(config.GetSecretKey())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func GetUsernameFromToken(authorization string) (string, error) {

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(authorization, claims, func(token *jwt.Token) (interface{}, error) {
		return config.GetSecretKey(), nil
	})

	if err != nil {
		log.Info(token)
		return "", err //errors.New("PostComment not autherized")
	}
	if token.Valid {
		username := claims["user"].(string)
		return username, nil
	}

	return "", errors.New("unable to extract UsernameFromToken")
}
