package auth

import (
	"database/sql"
	"errors"
	"fmt"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(name, passwordHash string) error {
	const query = `
		INSERT INTO users (username, password_hash)
		VALUES ($1, $2)
	`
	_, err := r.db.Exec(query, name, passwordHash)

	if err != nil {
		return err
	}

	return nil
}

func (r *postgresRepository) ReadHash(name string) (string, error) {
	const query = `
		SELECT password_hash FROM users WHERE username = $1
	`

	var hash string
	err := r.db.QueryRow(query, name).Scan(&hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user %q not found", name)
		}
		return "", err
	}

	return hash, nil
}