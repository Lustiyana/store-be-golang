package structs

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}