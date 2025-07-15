package repository

import (
	"context"
	"golang_test/internal/auth/entity"
)

type AuthRepository interface {
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}
