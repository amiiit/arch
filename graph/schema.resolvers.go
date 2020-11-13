package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"gitlab.com/amiiit/arco/graph/generated"
	"gitlab.com/amiiit/arco/graph/model"
)

func (r *mutationResolver) AddUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	// todo: authentication and roles
	user := input.ToUser()
	persistedUser, err := r.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	userModel := model.FromUser(persistedUser)
	return &userModel, nil
}

func (r *mutationResolver) EditUser(ctx context.Context, userID string, input model.UserInput) (*model.User, error) {
	userUpdate := input.ToUser()
	user, err := r.UserRepository.UpdateUser(ctx, userID, userUpdate)
	result := model.FromUser(user)

	return &result, err
}

func (r *mutationResolver) SetUserRoles(ctx context.Context, input model.SetRolesInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return []*model.User{}, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user, err := r.UserRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	result := model.FromUser(user)
	return &result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
