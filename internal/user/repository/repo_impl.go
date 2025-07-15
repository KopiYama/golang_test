package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"golang_test/internal/user/entity"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) GetAll(ctx context.Context) ([]entity.UserWithRole, error) {
	var users []entity.UserWithRole
	query := `
		SELECT 
			u.id, u.role_id, u.name, u.email, u.last_access,
			r.name as role_name
		FROM users u
		JOIN roles r ON r.id = u.role_id
	`
	err := r.db.SelectContext(ctx, &users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) Create(ctx context.Context, roleID int, name, email, password string) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO users (role_id, name, email, password)
		VALUES ($1, $2, $3, $4)
	`, roleID, name, email, password)
	return err
}

func (r *userRepo) Update(ctx context.Context, userID int, name string) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE users SET name = $1 WHERE id = $2
	`, name, userID)
	return err
}

func (r *userRepo) Delete(ctx context.Context, userID int) error {
	_, err := r.db.ExecContext(ctx, `
		DELETE FROM users WHERE id = $1
	`, userID)
	return err
}
