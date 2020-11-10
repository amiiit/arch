package auth

type Session struct {
	ID        string `db:"id" json:"id"`
	UserID    string `db:"user_id" json:"user_id"`
	Token     string `db:"hash"`
	CreatedAt string `db:"created_at" json:"created_at"`
	IsValid   bool   `db:"is_valid" json:"is_valid"`
}

type RoleType string
const AdminRole = RoleType("admin")

type Role struct {
	ID        string   `db:"id" json:"id"`
	UserID    string   `db:"user_id" json:"user_id"`
	CreatedAt string   `db:"created_at" json:"created_at"`
	Type      RoleType `db:"type" json:"type"`
}

type UserRoles struct {
	Admin bool
}
