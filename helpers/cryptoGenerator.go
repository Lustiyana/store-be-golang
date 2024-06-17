package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
  if err != nil {
		return "", err
  }

	return string(hash), nil
}

func ValidatePassword(hashPassword string, password string) (error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))

	if err != nil {
		return err
	}

	return nil
}