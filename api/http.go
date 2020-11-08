package api

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"gitlab.com/amiiit/arco/user"
	"net/http"
)

type createUserCmd struct {
	username  string
	firstName string
	lastName  string
	email     string
	password  string
}

type UserAPI struct {
	UserService    user.UserService
	UserRepository user.UserRepository
}

func (a UserAPI) HandleCreateUser(c echo.Context) error {
	cmd := &createUserCmd{}
	err := c.Bind(cmd)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	ctx := context.Background()
	newUser, err := a.UserService.CompleteUserObject(user.User{
		Email: cmd.email,
		FirstName: cmd.firstName,
		LastName: cmd.lastName,
	}, cmd.password)
	if err != nil {
		return c.String(http.StatusBadRequest, "validation failed")
	}

	persistedUser, err := a.UserRepository.CreateUser(ctx, newUser)
	if err != nil {
		if errors.Is(err, user.ErrUsernameTaken) {
			return c.String(http.StatusConflict, "username already exists")
		}
		return c.String(500, "error creating new user")
	}
	return c.JSON(http.StatusCreated, persistedUser)
}
