package user

import (
	"context"
	"errors"
	"fmt"
	sql "github.com/jmoiron/sqlx"
	"math/rand"
	"strings"
)

var ErrUsernameTaken = errors.New("username already taken")
var ErrUsernameNotFound = errors.New("user not found")
var tokenCharPool = "abcdefghijklmnopqrstuvwxyzABCEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func randomString(l int) string {
	bytes := make([]byte, l)

	for i := 0; i < l; i++ {
		bytes[i] = tokenCharPool[rand.Intn(len(tokenCharPool))]
	}
	return string(bytes)
}
func generateToken() string {
	return randomString(64)
}

type UserRepository struct {
	DB *sql.DB
}
type IUserRepository interface {
	CreateUser(ctx context.Context, user User) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	GetUserByID(ctx context.Context, username string) (User, error)
	CreateSession(ctx context.Context, userID string) (Session, error)
	GetSession(ctx context.Context, hash string) (Session, error)
	InvalidateSession(ctx context.Context, sessionID string) error
	GetRoles(ctx context.Context, userID string) (UserRoles, error)
	GetHashedPassword(ctx context.Context, userID string) (string, error)
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

func (r UserRepository) CreateSession(ctx context.Context, userID string) (Session, error) {
	s := Session{
		UserID:  userID,
		Token:   generateToken(),
		IsValid: true,
	}
	_, err := r.DB.NamedExecContext(ctx, `
		INSERT INTO sessions (
        	user_id, token              
		) VALUES (:user_id, :token)
	`, s)
	if err != nil {
		return Session{}, err
	}

	persistedSession, err := r.GetSession(ctx, s.Token)
	if err != nil {
		return Session{}, err
	}
	return persistedSession, nil
}

func (r UserRepository) GetSession(ctx context.Context, token string) (Session, error) {
	var session Session
	err := r.DB.GetContext(ctx, &session, `
		SELECT * FROM sessions WHERE token=$1
	`, token)
	if err != nil {
		return session, fmt.Errorf("error getting session via token: %w", err)
	}
	return session, nil
}

func (r UserRepository) InvalidateSession(ctx context.Context, sessionID string) error {
	_, err := r.DB.ExecContext(ctx, `
		UPDATE session SET is_valid = false WHERE id=$1
`, sessionID)
	return err
}

func (r UserRepository) GetRoles(ctx context.Context, userID string) (UserRoles, error) {
	var roles []Role
	err := r.DB.GetContext(ctx, &roles, `
		SELECT * FROM roles WHERE user_id=$1
`, userID)

	userRoles := UserRoles{}
	for _, role := range roles {
		switch role.Type {
		case "admin":
			userRoles.Admin = true
			break
		}
	}

	return userRoles, err
}

func (r UserRepository) GetHashedPassword(ctx context.Context, userID string) (string, error) {
	user := User{}
	err := r.DB.GetContext(ctx, &user, `
		SELECT hashed_password from users where id=$1
`, userID)
	if err != nil {
		return "", fmt.Errorf("retrieving hashed password failed: %w", err)
	}
	return user.HashedPassword, err
}
