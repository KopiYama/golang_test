package usecase

import (
	"context"
)

type AuthUsecase interface {
	Login(ctx context.Context, email, password string) (string, error)
	Logout(ctx context.Context, userID int) error
}
