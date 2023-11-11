package connect

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

const SecretKey = "15f5a4e7cfb1ec334f67433af0a4f86e563066945f7337314bd104a373fee500"

func VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, nil
}

func checkUser(token string) {
	valid, err := VerifyToken(token)
	if err != nil {
		fmt.Println("Token validation failed:", err)
	} else {
		fmt.Println("Token valid:", valid)
	}
}
