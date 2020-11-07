package user

import (
	"context"
	sql "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserRepository_CreateGetUser(t *testing.T) {
	conn, err := sql.Connect("postgres", "user=test dbname=arco_test sslmode=disable")
	require.NoError(t, err)
	repo := UserRepository{db: conn}
	user := User{
		Username:       "testuser",
		FirstName:      "Test",
		LastName:       "User",
		Email:          "test@user.net",
		Phone:          "+123456789",
		HashedPassword: "fbjslfuhew",
		PasswordSalt:   "123",
	}
	ctx := context.Background()
	err = repo.CreateUser(ctx, user)
	require.NoError(t, err)

	persistedUser, err := repo.GetUserByUsername(ctx, user.Username)
	require.NoError(t, err)
	require.Equal(t, user.Email, persistedUser.Email)

}
