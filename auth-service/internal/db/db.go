package db

import (
	"context"
)

type User struct {
	ID           string `db:"id"`
	Login        string `db:"login"`
	PasswordHash string `db:"password_hash"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, id, login, passwordHash string) error
	GetByLogin(ctx context.Context, login string) (*User, error)
	ExistsByLogin(ctx context.Context, login string) (bool, error)
}
