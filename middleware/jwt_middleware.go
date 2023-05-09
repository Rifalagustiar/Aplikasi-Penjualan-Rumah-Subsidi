package middleware

import (
	"crud/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userID uint, userName string, userType string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["userName"] = userName
	claims["userType"] = userType
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(constants.SECRET_JWT))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
