package verfication

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendEmailVerification(to, hash string) error {
	e := email.NewEmail()
	e.From = "Your Service <test@gmail.com>"
	e.To = []string{to}
	e.Subject = "Email Verification"
	e.HTML = []byte(fmt.Sprintf("<h1>Click the link to verify: <a href='http://localhost:8000/email/verify/%s'>Verify</a></h1>", hash))
	auth := smtp.PlainAuth("", "test@gmail.com", "1234", "smtp.gmail.com")
	return e.Send("smtp.gmail.com:587", auth)
}
