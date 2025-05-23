package auth

import (
	"context"
	"net/http"

	"github.com/TeddyMuli/go_graphql_api/internal/users"
	"github.com/TeddyMuli/go_graphql_api/internal/pkg/jwt"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		// Allow unauthenticated users in
		if header == "" {
			next.ServeHTTP(w, r)
			return
		}

		// Validate JWT token
		tokenStr := header
		username, err := jwt.ParseToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		// Fetch user from DB
		user := users.User{Username: username}
		id, err := users.GetUserIdByUsername(username)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		user.ID = id

		// Store user in context
		ctx := context.WithValue(r.Context(), userCtxKey, &user)
		r = r.WithContext(ctx)

		// Continue to next handler
		next.ServeHTTP(w, r)
	})
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}
