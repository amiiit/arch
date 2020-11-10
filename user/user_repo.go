package user

import (
	"context"
	"errors"
	"fmt"
	sql "github.com/jmoiron/sqlx"
	"strings"
)

var ErrUsernameTaken = errors.New("username already taken")

type UserRepository struct {
	DB *sql.DB
}
type IUserRepository interface {
	CreateUser(ctx context.Context, user User) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	GetUserByID(ctx context.Context, username string) (User, error)
}

func (rep UserRepository) CreateUser(ctx context.Context, user User) (User, error) {
	_, err := rep.DB.NamedExecContext(ctx, `
		INSERT INTO users (
			username, first_name, last_name, email, phone, region, hashed_password, password_salt
		) values (:username, :first_name, :last_name, :email, :phone, :region, :hashed_password, :password_salt)
`, user)
	if err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "users_username_key"`) {
			return User{}, ErrUsernameTaken
		}
		return User{}, err
	}
	return rep.GetUserByUsername(ctx, user.Username)
}

func (rep UserRepository) GetUserByUsername(ctx context.Context, username string) (User, error) {
	var user User
	err := rep.DB.GetContext(ctx, &user, `SELECT * from users where username = $1`, username)
	if err != nil {
		return user, fmt.Errorf("error getting user by name: %w", err)
	}
	return user, err
}

func (rep UserRepository) GetUserByID(ctx context.Context, id string) (User, error) {
	var user User
	err := rep.DB.GetContext(ctx, &user, `SELECT * from users where id = $1`, id)
	if err != nil {
		return user, fmt.Errorf("error getting user by id: %w", err)
	}
	return user, err
}
