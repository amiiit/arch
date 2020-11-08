package api

import (
	"gitlab.com/amiiit/arco/user"
	"testing"
)


func TestUserAPI_HandleCreateUser(t *testing.T) {
	userAPI := UserAPI{
		UserService:    user.UserService{},
		UserRepository: user.UserRepository{},
	}
}