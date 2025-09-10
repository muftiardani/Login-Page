package postgres

import (
	"context"
	"fmt"
	"log"
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

// GetUser mengambil data pengguna dari database berdasarkan email.
func (s *PostgresUserStore) GetUser(email string) (model.User, bool) {
	var user model.User
	query := "SELECT email, password_hash FROM users WHERE email = $1"

	err := s.DB.QueryRow(context.Background(), query, email).Scan(&user.Email, &user.PasswordHash)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.User{}, false
		}
		log.Printf("ERROR: Terjadi kesalahan saat mengambil data pengguna '%s': %v", email, err)
		return model.User{}, false
	}

	return user, true
}

// CreateUser memasukkan data pengguna baru ke dalam database.
func (s *PostgresUserStore) CreateUser(user model.User) error {
	query := "INSERT INTO users (email, password_hash) VALUES ($1, $2)"

	_, err := s.DB.Exec(context.Background(), query, user.Email, user.PasswordHash)
	if err != nil {
		return fmt.Errorf("kesalahan saat menyimpan pengguna ke database: %w", err)
	}

	return nil
}

// UpdateUser memperbarui data pengguna di database.
func (s *PostgresUserStore) UpdateUser(oldEmail string, user model.User) error {
	query := "UPDATE users SET email = $1, password_hash = $2 WHERE email = $3"

	_, err := s.DB.Exec(context.Background(), query, user.Email, user.PasswordHash, oldEmail)
	if err != nil {
		return fmt.Errorf("kesalahan saat memperbarui pengguna di database: %w", err)
	}

	return nil
}