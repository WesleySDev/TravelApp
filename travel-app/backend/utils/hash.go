package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GeraHashSenha(Password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparaSenha(Password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(Password))
}
