package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"golang_test/internal/auth/entity"
)

type authRepo struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) AuthRepository {
	return &authRepo{db: db}
}

func (r *authRepo) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.db.GetContext(ctx, &user, `SELECT * FROM users WHERE email = $1 LIMIT 1`, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
