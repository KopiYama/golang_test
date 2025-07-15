package usecase

import (
	"context"
	"errors"
	"golang_test/internal/auth/repository"
)

type authUsecase struct {
	repo repository.AuthRepository
}

func NewAuthUsecase(r repository.AuthRepository) AuthUsecase {
	return &authUsecase{repo: r}
}

func (u *authUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if user.Password != password {
		return "", errors.New("invalid credentials")
	}

	// sementara return token dummy
	return "isi random token", nil
}
