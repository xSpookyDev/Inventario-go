package entity

type UserRole struct {
	ID     int `db:"id"`
	UserID int `db:"user_id"`
	RoleID int `db:"role_id"`
}
