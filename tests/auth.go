package tests

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func GenerateToken(identity string, id uint) string {
	var key = []byte(os.Getenv("JTW_SECRET_KEY"))
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := token.Claims.(jwt.MapClaims)
	claims["Id"] = id
	claims["Email"] = identity
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["orig_iat"] = time.Now().Unix()
	tokenString, _ := token.SignedString(key)

	return tokenString
}