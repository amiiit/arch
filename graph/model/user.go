package model

import (
	"gitlab.com/amiiit/arco/user"
)

func FromUser(u user.User) User {
	return User{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     &u.Phone,
		Region:    &u.Region,
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
