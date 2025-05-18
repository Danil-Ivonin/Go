package repository

import (
	"fmt"
	todo "github.com/Danil-Ivonin/Go"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id;", usersTable)
	raw := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := raw.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
