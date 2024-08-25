package repository

import (
	"fmt"

	"github.com/co1seam/Taskify/service/backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash, email, birthdate, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, %6, $7) RETURNING id", user)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password, user.Email, user.Birthdate, user.CreatedAt, user.UpdatedAt)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}