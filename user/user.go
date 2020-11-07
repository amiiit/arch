package user

import "time"

type User struct {
	Username       string    `db:"username"`
	FirstName      string    `db:"first_name"`
	LastName       string    `db:"last_name"`
	Email          string    `db:"email"`
	Phone          string    `db:"phone"`
	HashedPassword string    `db:"hashed_password"`
	PasswordSalt   string    `db:"password_salt"`
	CreatedAt      time.Time `db:"created_at"`
	LastUpdated    time.Time `db:"last_updated"`
}
