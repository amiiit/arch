package api

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"gitlab.com/amiiit/arco/mocks"
	"gitlab.com/amiiit/arco/user"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserAPI_HandleCreateUser(t *testing.T) {
	userService := mocks.NewIUserServiceMock(t)
	userRepository := mocks.NewIUserRepositoryMock(t)

	userAPI := UserAPI{
		UserService:    userService,
		UserRepository: userRepository,
	}

	cmd := createUserCmd{
		Username:  "testuser",
		FirstName: "firstname",
		LastName:  "lastname",
		Email:     "Email@me.com",
		Password:  "pa33wor∂",
	}
	mockUserToPersist := user.User{
		Username: cmd.Username,
		Email: cmd.Email,
		FirstName: cmd.FirstName,
		LastName: cmd.LastName,
	}
	userService.SetUserPasswordMock.Inspect(func(user user.User, password string) {
		require.Equal(t, cmd.Username, user.Username)
		require.Equal(t, cmd.Password, password)
	}).Return(mockUserToPersist, nil)

	userRepository.CreateUserMock.Inspect(func(ctx context.Context, createUserUser user.User) {
		require.Equal(t, mockUserToPersist, createUserUser)
	}).Return(user.User{}, nil)

	e := echo.New()
	payload, err := json.Marshal(cmd)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(payload)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = userAPI.HandleCreateUser(c)
	require.NoError(t, err)
}
