package user

import (
	"context"
	"net/http"
)

type contextKey string

const SessionContextKey = contextKey("session")
const RolesContextKey = contextKey("roles")

// Middleware decodes the share session cookie and packs the session into context
func Middleware(userRepo IUserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			token := r.Header.Get("auth")

			// Allow unauthenticated users in
			if token == "" {
				next.ServeHTTP(w, r)
				return
			}

			session, err := userRepo.GetSession(ctx, token)
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

			roles, err := userRepo.GetRoles(ctx, session.UserID)
			if err != nil {
				http.Error(w, "Error fetching roles", http.StatusInternalServerError)
				return
			}
			ctx = context.WithValue(ctx, RolesContextKey, roles)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
