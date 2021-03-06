package graph

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"gitlab.com/amiiit/arco/graph/generated"
	"gitlab.com/amiiit/arco/user"
)

type Directives struct {
}

func (d Directives) Apply(root *generated.DirectiveRoot) {
	root.HasRole = d.HasRole
}

func (d Directives) HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role string) (res interface{}, err error) {
	contextRoles := ctx.Value(user.RolesContextKey)
	if contextRoles == nil {
		return nil, fmt.Errorf("Access denied, please log in.")
	}
	roles := contextRoles.(user.UserRoles)
	allow := false

	switch role {
	case "admin":
		allow = roles.Admin
		break
	case "member":
		allow = roles.Member
		break
	}

	if !allow {
		return nil, fmt.Errorf("access denied, must be %s", role)
	}

	return next(ctx)
}
