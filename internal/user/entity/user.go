package entity

import "time"

type User struct {
	ID         int       `db:"id"`
	RoleID     int       `db:"role_id"`
	Name       string    `db:"name"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	LastAccess time.Time `db:"last_access"`
}

type UserWithRole struct {
	ID         int       `db:"id"`
	RoleID     int       `db:"role_id"`
	RoleName   string    `db:"role_name"`
	Name       string    `db:"name"`
	Email      string    `db:"email"`
	LastAccess time.Time `db:"last_access"`
}
