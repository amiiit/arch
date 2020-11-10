package api

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"gitlab.com/amiiit/arco/user"
	"net/http"
)

type createUserCmd struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserAPI struct {
	UserService    user.IUserService
	UserRepository user.IUserRepository
}

func (a UserAPI) HandleCreateUser(c echo.Context) error {
	cmd := &createUserCmd{}
	err := c.Bind(cmd)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	ctx := context.Background()
	newUser, err := a.UserService.SetUserPassword(user.User{
		Username: cmd.Username,
		Email: cmd.Email,
		FirstName: cmd.FirstName,
		LastName: cmd.LastName,
	}, cmd.Password)
	if err != nil {
		return c.String(http.StatusBadRequest, "validation failed")
	}

	persistedUser, err := a.UserRepository.CreateUser(ctx, newUser)
	if err != nil {
		if errors.Is(err, user.ErrUsernameTaken) {
			return c.String(http.StatusConflict, "Username already exists")
		}
		return c.String(500, "error creating new user")
	}
	return c.JSON(http.StatusCreated, persistedUser)
}
