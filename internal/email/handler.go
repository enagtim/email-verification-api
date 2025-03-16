package email

import (
	"api/email-verification/pkg/verfication"
	"encoding/json"
	"net/http"
)

type EmailHandler struct {
	EmailRepository *EmailRepository
}

func NewEmailHandler(router *http.ServeMux, deps EmailHandler) {
	handler := &EmailHandler{
		EmailRepository: deps.EmailRepository,
	}
	router.HandleFunc("POST /email", handler.Create())
	router.HandleFunc("GET /email/verify/{hash}", handler.Verify())
}

func (handler *EmailHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req EmailRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "ERROR PARSING JSON", http.StatusBadRequest)
			return
		}
		email := NewEmail(req.Email)
		err = handler.EmailRepository.Create(email)
		if err != nil {
			http.Error(w, "ERROR SAVE DATABASE", http.StatusInternalServerError)
			return
		}
		err = verfication.SendEmailVerification(email.Email, email.Hash)
		if err != nil {
			http.Error(w, "ERROR SENDING EMAIL", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
	}
}
func (handler *EmailHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		susccess, err := handler.EmailRepository.Verify(hash)
		if err != nil {
			http.Error(w, "ERROR VERIFYING HASH", http.StatusInternalServerError)
			return
		}
		if !susccess {
			http.Error(w, "INVALID HASH", http.StatusNotFound)
			return
		}
		w.WriteHeader(200)
	}
}
