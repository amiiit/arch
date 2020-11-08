package user

import (
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
	HashedPassword string    `db:"hashed_password"`
	PasswordSalt   string    `db:"password_salt"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	LastUpdated    time.Time `db:"last_updated" json:"last_updated"`
}

type UserService struct {
	repo UserRepository
}

type IUserService interface {
	CompleteUserObject(user User, password string) (User, error)
}

func (s UserService) CompleteUserObject(user User, password string) (User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.HashedPassword = string(hashed)
	// todo: validation

	return user, nil
}
