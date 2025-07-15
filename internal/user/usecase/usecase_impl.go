package usecase

import (
	"context"
	"golang_test/internal/user/entity"
	"golang_test/internal/user/repository"
)

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: repo}
}

func (u *userUsecase) GetAll(ctx context.Context) ([]entity.UserWithRole, error) {
	return u.userRepo.GetAll(ctx)
}

func (u *userUsecase) Create(ctx context.Context, roleID int, name, email, password string) error {
	return u.userRepo.Create(ctx, roleID, name, email, password)
}

func (u *userUsecase) Update(ctx context.Context, userID int, name string) error {
	return u.userRepo.Update(ctx, userID, name)
}

func (u *userUsecase) Delete(ctx context.Context, userID int) error {
	return u.userRepo.Delete(ctx, userID)
}
