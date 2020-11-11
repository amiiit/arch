package auth

import (
	"context"
	"fmt"
	"gitlab.com/amiiit/arco/user"
	"net/http"
)

type contextKey string
const SessionContextKey = contextKey("session")
const RolesContextKey = contextKey("roles")

// Middleware decodes the share session cookie and packs the session into context
func Middleware(authRepo IAuthRepository, userRepo user.IUserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			token, err := r.Cookie("auth-token")

			// Allow unauthenticated users in
			if err != nil || token == nil {
				next.ServeHTTP(w, r)
				return
			}

			session, err := authRepo.GetSession(ctx, token.Value)
			if err != nil {
				http.Error(w, "Invalid cookie", http.StatusForbidden)
				return
			}

			// get the user from the database
			sessionUser, err := userRepo.GetUserByID(ctx, session.UserID)
			if err != nil {
				http.Error(w, "Error fetching user", http.StatusInternalServerError)
				return
			}
			ctx = context.WithValue(ctx, SessionContextKey, sessionUser)

			roles, err := authRepo.GetRoles(ctx, session.UserID)
			if err != nil {
				http.Error(w, "Error fetching roles", http.StatusInternalServerError)
				return
			}
			fmt.Println("middleware roles", roles)
			ctx = context.WithValue(ctx, RolesContextKey, roles)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
