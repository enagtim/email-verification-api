package email

import "math/rand"

type Email struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

func NewEmail(email string) *Email {
	return &Email{
		Email: email,
		Hash:  RandStringRunes(6),
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for index := range b {
		b[index] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
