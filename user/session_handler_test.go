package user

import (
	"context"
	sql "github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestSessionHandler_CreateSession(t *testing.T) {
	conn, err := sql.Connect("postgres", "user=test dbname=arco_test sslmode=disable")
	_, err = conn.Exec("delete from users;")
	require.NoError(t, err)
	repo := UserRepository{DB: conn}
	service := UserService{Repo: repo}
	ctx := context.Background()
	password := "p433w0r∂"
	user := User{
		Username:  "sessiontestuser",
		FirstName: "Session",
		LastName:  "Test",
		Email:     "session@te.st",
	}
	user, err = service.SetUserPassword(user, password)
	require.NoError(t, err)

	user, err = repo.CreateUser(ctx, user)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/session", nil)
	require.NoError(t, err)
	req.Form = url.Values{}
	req.Form.Set("username", user.Username)
	req.Form.Set("password", password)

	sessionHandler := SessionHandler{userService: service}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(sessionHandler.CreateSession)
	handler.ServeHTTP(rr, req)

	require.Equal(t, "session created", rr.Body.String())
}
