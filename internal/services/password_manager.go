package services

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PasswordManager interface {
	GenerateSalt() (string, error)
	HashPassword(password, salt string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type PasswordManagerImpl struct{}

func (pass *PasswordManagerImpl) GenerateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

func (pass *PasswordManagerImpl) HashPassword(password, salt string) (string, error) {
	combined := password + salt
	if len([]byte(combined)) > 72 {
		return "", fmt.Errorf("password is too long")
	}
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(combined), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

func (pass *PasswordManagerImpl) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
