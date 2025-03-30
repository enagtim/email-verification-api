package email

import (
	"api/email-verification/pkg/db"
	"database/sql"
	"fmt"
)

type EmailRepository struct {
	Database *db.Db
}

func NewEmailRepository(database *db.Db) *EmailRepository {
	return &EmailRepository{
		Database: database,
	}
}
func (repo *EmailRepository) Create(email *Email) error {
	query := `INSERT INTO emails (email, hash) VALUES($1, $2)`
	_, err := repo.Database.Exec(query, email.Email, email.Hash)
	if err != nil {
		return fmt.Errorf("ERROR SAVE EMAIL: %v", err)
	}
	return nil
}

func (repo *EmailRepository) Verify(hash string) (bool, error) {
	var email string
	query := `SELECT email FROM emails WHERE hash = $1`
	err := repo.Database.QueryRow(query, hash).Scan(&email)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("ERROR CHECK HASH: %v", err)
	}
	deleteQuery := `DELETE FROM emails WHERE hash = $1`
	_, err = repo.Database.Exec(deleteQuery, hash)
	if err != nil {
		return false, fmt.Errorf("ERROR DELETED: %v", err)
	}

	return true, nil

}
