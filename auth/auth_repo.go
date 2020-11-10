package auth

import (
	"context"
	"fmt"
	sql "github.com/jmoiron/sqlx"
	"math/rand"
)

var tokenCharPool = "abcdefghijklmnopqrstuvwxyzABCEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func randomString(l int) string {
	bytes := make([]byte, l)

	for i := 0; i < l; i++ {
		bytes[i] = tokenCharPool[rand.Intn(len(tokenCharPool))]
	}
	return string(bytes)
}
func generateToken() string {
	return randomString(64)
}

type IAuthRepository interface {
	CreateSession(ctx context.Context, userID string) (Session, error)
	GetSession(ctx context.Context, hash string) (Session, error)
	InvalidateSession(ctx context.Context, sessionID string) error
	GetRoles(ctx context.Context, userID string) (UserRoles, error)
}

type Repository struct {
	DB *sql.DB
}

func (r Repository) CreateSession(ctx context.Context, userID string) (Session, error) {
	s := Session{
		UserID:  userID,
		Token:   generateToken(),
		IsValid: true,
	}
	_, err := r.DB.NamedExecContext(ctx, `
		INSERT INTO sessions (
        	user_id, token              
		) VALUES (:user_id, :token)
	`, s)
	if err != nil {
		return Session{}, err
	}

	persistedSession, err := r.GetSession(ctx, s.Token)
	if err != nil {
		return Session{}, err
	}
	return persistedSession, nil
}

func (r Repository) GetSession(ctx context.Context, token string) (Session, error) {
	var session Session
	err := r.DB.GetContext(ctx, &session, `
		SELECT * FROM sessions WHERE token=$1
	`, token)
	if err != nil {
		return session, fmt.Errorf("error getting session via token: %w", err)
	}
	return session, nil
}

func (r Repository) InvalidateSession(ctx context.Context, sessionID string) error {
	_, err := r.DB.ExecContext(ctx, `
		UPDATE session SET is_valid = false WHERE id=$1
`, sessionID)
	return err
}

func (r Repository) GetRoles(ctx context.Context, userID string) (UserRoles, error) {
	var roles []Role
	err := r.DB.GetContext(ctx, &roles, `
		SELECT * FROM roles WHERE user_id=$1
`, userID)

	userRoles := UserRoles{}
	for _, role := range roles {
		switch role.Type {
		case "admin":
			userRoles.Admin = true
			break
		}
	}

	return userRoles, err
}
