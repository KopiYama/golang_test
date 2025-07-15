package usecase

import (
	"context"
)

type AuthUsecase interface {
	Login(ctx context.Context, email, password string) (string, error)
}
