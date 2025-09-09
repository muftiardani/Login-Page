package postgres

import (
	"context"
	"fmt"
	"login-api/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserStore struct {
	DB *pgxpool.Pool
}

func NewPostgresUserStore(db *pgxpool.Pool) *PostgresUserStore {
	return &PostgresUserStore{DB: db}
}

func (s *PostgresUserStore) GetUser(email string) (model.User, bool) {
	var user model.User
	query := "SELECT email, password_hash FROM users WHERE email = $1"

	err := s.DB.QueryRow(context.Background(), query, email).Scan(&user.Email, &user.PasswordHash)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.User{}, false
		}
		return model.User{}, false
	}

	return user, true
}

func (s *PostgresUserStore) CreateUser(user model.User) error {
	query := "INSERT INTO users (email, password_hash) VALUES ($1, $2)"

	_, err := s.DB.Exec(context.Background(), query, user.Email, user.PasswordHash)
	if err != nil {
		return fmt.Errorf("database error on create user: %w", err)
	}

	return nil
}