package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"golang_test/internal/auth/repository"
	"golang_test/internal/shared/jwt"
	"golang_test/internal/shared/redis"
	"time"
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
		return "", err
	}

	token, err := jwt.GenerateToken(user.ID, user.Name)
	if err != nil {
		return "", err
	}

	// marshal user ke JSON
	userJson, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	// simpan user ke Redis selama 1 jam
	err = redis.Client.Set(ctx, fmt.Sprintf("auth:user:%d", user.ID), userJson, time.Hour*1).Err()
	if err != nil {
		return "", err
	}

	return token, nil
}
