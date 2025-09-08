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

// GetUser mencari pengguna berdasarkan username di database.
func (s *PostgresUserStore) GetUser(username string) (model.User, bool) {
	var user model.User
	query := "SELECT username, password_hash FROM users WHERE username = $1"

	err := s.DB.QueryRow(context.Background(), query, username).Scan(&user.Username, &user.PasswordHash)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.User{}, false
		}
		return model.User{}, false
	}

	return user, true
}

// CreateUser menyimpan pengguna baru ke dalam database.
func (s *PostgresUserStore) CreateUser(user model.User) error {
	query := "INSERT INTO users (username, password_hash) VALUES ($1, $2)"

	_, err := s.DB.Exec(context.Background(), query, user.Username, user.PasswordHash)
	if err != nil {
		return fmt.Errorf("database error on create user: %w", err)
	}

	return nil
}