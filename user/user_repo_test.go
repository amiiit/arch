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
	_, err = conn.Exec("delete from users;")
	require.NoError(t, err)
	repo := UserRepository{DB: conn}

	t.Run("Create and retrieve user", func(t *testing.T) {
		user := User{
			Username:       "testuser",
			FirstName:      "Test",
			LastName:       "Member",
			Email:          "test@user.net",
			Phone:          "+123456789",
			HashedPassword: "fbjslfuhew",
		}
		ctx := context.Background()
		_, err = repo.CreateUser(ctx, user)
		require.NoError(t, err)

		persistedUser, err := repo.GetUserByUsername(ctx, user.Username)
		require.NoError(t, err)
		require.NotNil(t, persistedUser.ID)
		require.Equal(t, user.Email, persistedUser.Email)
	})

	t.Run("Create dupe username", func(t *testing.T) {
		user := User{
			Username:       "mustbeunique",
			FirstName:      "Test",
			LastName:       "Member",
			Email:          "test@user1.net",
			Phone:          "+123456789",
			HashedPassword: "fbjslfuhew",
		}
		ctx := context.Background()
		pUser, err := repo.CreateUser(ctx, user)
		require.NoError(t, err)
		require.NotNil(t, pUser.ID)

		user.Email = "another@mail.com"
		user.Phone = "+987654321"

		_, dupeErr := repo.CreateUser(ctx, user)
		require.EqualError(t, dupeErr, ErrUsernameTaken.Error())
	})

	t.Run("Update user", func(t *testing.T) {
		ctx := context.Background()
		testUser, err := repo.GetUserByUsername(ctx, "testuser")
		require.NoError(t, err)
		require.NotEmpty(t, testUser.Email)
		oldEmail := testUser.Email
		userID := testUser.ID

		newPhoneNumber := "+34678678678"
		require.NotEqual(t, newPhoneNumber, testUser)
		testUser.Phone = newPhoneNumber
		testUser.ID = ""

		updatedUser, err := repo.UpdateUser(ctx, userID, testUser)
		require.NoError(t, err)
		require.Equal(t, newPhoneNumber, updatedUser.Phone)
		require.Equal(t, newPhoneNumber, updatedUser.Phone)
		require.Equal(t, oldEmail, updatedUser.Email)
	})
}