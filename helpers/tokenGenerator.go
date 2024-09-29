package helpers

import (
	"fmt"
	"time"
	"net/http"
	"strings"
	"store-be-golang/structs"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("1NIP45SW0RD")

func GenerateToken(id uint, email string, password string) (string, error) {
	claims := &structs.Claims{
		ID: id,
		Email: email,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
			IssuedAt:  time.Now().Unix(),
			Issuer:    "golang_sanbercode",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*structs.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &structs.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*structs.Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return claims, nil
}


func ExtractToken(auth string) (string, error) {
	token := strings.TrimPrefix(auth, "Bearer ")
	if token == auth {
		return "", http.ErrNotSupported
	}
	return token, nil
}