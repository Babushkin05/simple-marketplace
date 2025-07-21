package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type PostgresRepo struct {
	DB *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{DB: db}
}

func (r *PostgresRepo) CreateUser(ctx context.Context, id, login, passwordHash string) error {
	query := `
		INSERT INTO users (id, login, password_hash)
		VALUES ($1, $2, $3)
	`
	_, err := r.DB.ExecContext(ctx, query, id, login, passwordHash)
	return err
}

func (r *PostgresRepo) GetByLogin(ctx context.Context, login string) (*User, error) {
	var u User
	query := `
		SELECT id, login, password_hash
		FROM users
		WHERE login = $1
	`
	err := r.DB.GetContext(ctx, &u, query, login)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("user not found")
	}
	return &u, err
}

func (r *PostgresRepo) ExistsByLogin(ctx context.Context, login string) (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1 FROM users WHERE login = $1
		)
	`
	err := r.DB.GetContext(ctx, &exists, query, login)
	return exists, err
}
