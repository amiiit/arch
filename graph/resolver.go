package graph

import "gitlab.com/amiiit/arco/user"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	UserService user.IUserService
	UserRepository user.IUserRepository
}
