package controllers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/moroz/lenslocked/models"
)

func AuthenticateUserMiddleware(db *sql.DB, cookieKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(cookieKey)
			if err != nil {
				ctx := context.WithValue(r.Context(), "user", nil)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			token := cookie.Value
			user := models.AuthenticateUserByToken(db, token)
			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RestrictAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawUser := r.Context().Value("user")
		if rawUser == nil {
			http.Redirect(w, r, "/sign-in", http.StatusFound)
			return
		}
		user := rawUser.(*models.User)
		if user == nil {
			http.Redirect(w, r, "/sign-in", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
