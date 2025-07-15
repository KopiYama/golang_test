package entity

type User struct {
	ID         int    `db:"id"`
	RoleID     int    `db:"role_id"`
	Name       string `db:"name"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	LastAccess string `db:"last_access"`
}
