package user

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID             string    `db:"id" json:"id"`
	Username       string    `db:"username" json:"username"`
	FirstName      string    `db:"first_name" json:"first_name"`
	LastName       string    `db:"last_name" json:"last_name"`
	Email          string    `db:"email" json:"email"`
	Phone          string    `db:"phone" json:"phone"`
	Region         string    `db:"region" json:"region"`
	HashedPassword string    `db:"hashed_password"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	LastUpdated    time.Time `db:"last_updated" json:"last_updated"`
}

type UserService struct {
	Repo IUserRepository
}

type IUserService interface {
	SetUserPassword(user User, password string) (User, error)
	ValidatePassword(ctx context.Context, username string, password string) error
	CreateSession(ctx context.Context, username string) (Session, error)
}

func (s UserService) SetUserPassword(user User, password string) (User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.HashedPassword = string(hashed)
	// todo: validation

	return user, nil
}

func (s UserService) ValidatePassword(ctx context.Context, username string, password string) error {
	user, err := s.Repo.GetUserByUsername(ctx, username)
	if err != nil {
		return fmt.Errorf("failed fetching user: %w", err)
	}
	if user.ID == "" {
		return ErrUsernameNotFound
	}
	hashed, err := s.Repo.GetHashedPassword(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("failed getting hased password: %w", err)
	}
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}

func (s UserService) CreateSession(ctx context.Context, username string) (Session, error) {
	user, err := s.Repo.GetUserByUsername(ctx, username)
	if err != nil {
		return Session{}, fmt.Errorf("failed fetching user: %w", err)
	}
	return s.Repo.CreateSession(ctx, user.ID)
}