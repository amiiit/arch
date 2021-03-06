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
		Suspended: u.Suspended,
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

func FromUserRoles(r user.UserRoles) UserRoles {
	return UserRoles{
		Admin:  r.Admin,
		Member: r.Member,
	}
}