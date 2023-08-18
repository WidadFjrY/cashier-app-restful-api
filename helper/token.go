package helper

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("sjhdjhihdoahssjcbabduiaghdwuiuah928319038jasfhji1289y39jkashdkj"))
}
