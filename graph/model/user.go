package model

import (
	"gitlab.com/amiiit/arco/user"
)

func FromUser(um user.User) User {
	return User{
		ID:        um.ID,
		Username:  um.Username,
		FirstName: um.FirstName,
		LastName:  um.LastName,
		Email:     um.Email,
		Phone:     &um.Phone,
		Region:    &um.Region,
	}
}
func (ui UserInput) ToUser() user.User {
	return user.User{
		Username:  ui.Username,
		FirstName: ui.FirstName,
		LastName:  ui.LastName,
		Email:     ui.Email,
		Phone:     *ui.Phone,
		Region:    *ui.Region,
	}
}
