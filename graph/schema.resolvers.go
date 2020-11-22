package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"gitlab.com/amiiit/arco/graph/generated"
	"gitlab.com/amiiit/arco/graph/model"
	"gitlab.com/amiiit/arco/user"
)

func (r *mutationResolver) SetOfferDescription(ctx context.Context, offerID string, input *model.OfferDescriptionInput) (*model.Offer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveOfferDescription(ctx context.Context, input model.DeleteOfferDescriptionInput) (*model.Offer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddOffer(ctx context.Context, input model.OfferInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditOffer(ctx context.Context, offerID string, input model.OfferInput) (*model.Offer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RequestOfferRevision(ctx context.Context, offerID string) (*model.Offer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SetOfferPublishedState(ctx context.Context, offerID string, state *model.PublishedState) (*model.Offer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	// todo: authentication and roles
	userInput := input.ToUser()
	persistedUser, err := r.UserRepository.CreateUser(ctx, userInput)
	if err != nil {
		return nil, err
	}
	userModel := model.FromUser(persistedUser)
	return &userModel, nil
}

func (r *mutationResolver) SetUserPassword(ctx context.Context, userID string, newPassword string) (*model.User, error) {
	user, err := r.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	user, err = r.UserService.SetUserPassword(user, newPassword)
	if err != nil {
		return nil, err
	}
	updatedUser, err := r.UserRepository.UpdateUser(ctx, userID, user)
	if err != nil {
		return nil, err
	}

	result := model.FromUser(updatedUser)
	return &result, nil
}

func (r *mutationResolver) EditUser(ctx context.Context, userID string, input model.UserInput) (*model.User, error) {
	userUpdate := input.ToUser()
	user, err := r.UserRepository.UpdateUser(ctx, userID, userUpdate)
	result := model.FromUser(user)

	return &result, err
}

func (r *mutationResolver) SetUserRoles(ctx context.Context, userID string, roles model.RolesInput) (*model.User, error) {
	err := r.UserRepository.SetUserRoles(ctx, userID, user.UserRoles{
		Admin:  roles.Admin,
		Member: roles.Member,
		UserID: userID,
	})

	persistedUser, err := r.UserRepository.GetUserByID(ctx, userID)
	userModel := model.FromUser(persistedUser)
	return &userModel, err
}

func (r *mutationResolver) SetUserSuspended(ctx context.Context, userID string, suspended bool) (*model.User, error) {
	updatedUser, err := r.UserRepository.SetUserSuspended(ctx, userID, suspended)
	if err != nil {
		return nil, err
	}
	result := model.FromUser(updatedUser)
	return &result, err
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	contextUser := ctx.Value(user.UserContextKey)
	if contextUser == nil {
		return nil, fmt.Errorf("session is not assigned to a user")
	}
	sessionUser := contextUser.(user.User)
	result := model.FromUser(sessionUser)
	return &result, nil
}

func (r *queryResolver) Users(ctx context.Context, pagination *model.Pagination) ([]*model.User, error) {
	users, err := r.UserRepository.GetUsers(ctx, user.Pagination{
		Limit:  pagination.Limit,
		Offset: pagination.Offset,
	})
	if err != nil {
		return nil, fmt.Errorf("fetching users failed: %w", err)
	}

	result := make([]*model.User, len(users))
	for i, user := range users {
		userModel := model.FromUser(user)
		result[i] = &userModel
	}
	return result, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user, err := r.UserRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	result := model.FromUser(user)
	return &result, nil
}

func (r *userResolver) Roles(ctx context.Context, obj *model.User) (*model.UserRoles, error) {
	roles, err := r.UserRepository.GetRoles(ctx, obj.ID)
	if err != nil {
		return nil, fmt.Errorf("getting roles for user failed: %w", err)
	}
	modelRoles := model.FromUserRoles(roles)
	return &modelRoles, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
