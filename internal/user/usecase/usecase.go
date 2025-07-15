package usecase

import (
	"context"
	"golang_test/internal/user/entity"
)

type UserUsecase interface {
	GetAll(ctx context.Context) ([]entity.UserWithRole, error)
	Create(ctx context.Context, roleID int, name, email, password string) error
	Update(ctx context.Context, userID int, name string) error
	Delete(ctx context.Context, userID string) error
}
