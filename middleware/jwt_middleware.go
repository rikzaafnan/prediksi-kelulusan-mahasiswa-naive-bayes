package middleware

import (
	"time"

	"prediksi-kelulusan-mahasiswa-naive-bayes/constants"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId int, name string) (string, error) {

	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	// claims["exp"] = time.Now().Minute()
	// claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	// claims["liat-expired"] = time.Now().Add(time.Minute * 1)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SECRET_JWT))

}
