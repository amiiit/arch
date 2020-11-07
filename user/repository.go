package user

import (
	"context"
	"fmt"
	sql "github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sql.DB
}

func (rep UserRepository) CreateUser(ctx context.Context, user User) error {
	_, err := rep.db.NamedExecContext(ctx, `
		INSERT INTO users (
			username, first_name, last_name, email, phone, hashed_password, password_salt
		) values (:username, :first_name, :last_name, :email, :phone, :hashed_password, :password_salt)
`, user)
	return err
}

func (rep UserRepository) GetUserByUsername(ctx context.Context, username string) (User, error) {
	var user User
	err := rep.db.Get(&user, `SELECT * from users where username = $1`, username)
	if err != nil {
		return user, fmt.Errorf("error getting user by name: %w", err)
	}
	return user, err
}
