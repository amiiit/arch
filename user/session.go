package user

import (
	"net/http"
)

type SessionHandler struct {
	UserService IUserService
}

// CreateSession starts a new session via a POST request with `password` and `username`
// being sent as form-data.
// curl -X POST -F 'username=amiiit' -F 'password=amit' http://localhost:8080/session
func (h SessionHandler) CreateSession(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad form"))
		return
	}
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	if len(password) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("password missing"))
		return
	}

	if len(username) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("must provide either username or email"))
		return
	}
	validationError := h.UserService.ValidatePassword(ctx, username, password)
	if validationError != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("validation failed"))
		return
	}
	session, err := h.UserService.CreateSession(ctx, username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("creating new session failed"))
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:       "auth-token",
		Value:      session.Token,
		Secure:     false,
		HttpOnly:   false,
	})
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("session created"))
}

// InvalidateSession terminates the session associated
func (h SessionHandler) InvalidateSession(writer http.ResponseWriter, request *http.Request) {

}
