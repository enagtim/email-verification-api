package email

import (
	"crypto/rand"
	"encoding/base64"
)

type Email struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

func NewEmail(email string) *Email {
	hash, err := GenerateVerificationHash()
	if err != nil {
		return nil
	}
	return &Email{
		Email: email,
		Hash:  hash,
	}
}
func GenerateVerificationHash() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
