package main

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/cisco100/wepost/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
)

func (app *Application) AuthTokenMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")

			if authHeader == " " {
				app.UnauthorizedError(w, r, fmt.Errorf("authorization header is missing"))
				return
			}

			parts := strings.Split(authHeader, " ")

			if len(parts) != 2 || parts[0] != "Bearer" {
				app.UnauthorizedBasicAuthError(w, r, fmt.Errorf("authorization header is malformed"))
				return
			}

			token := parts[1]

			jwtToken, err := app.Authenticator.ValidateToken(token)
			if err != nil {
				app.UnauthorizedError(w, r, err)
				return
			}

			claims, _ := jwtToken.Claims.(jwt.MapClaims)

			userID := claims["subs"].(string)

			ctx := r.Context()
			user, err := app.Store.User.GetUserById(ctx, userID)
			if err != nil {
				app.UnauthorizedError(w, r, err)
				return
			}

			ctx = context.WithValue(ctx, userCtx, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (app *Application) PostsContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "postID")

		ctx := r.Context()

		post, err := app.Store.Post.GetPostById(ctx, idParam)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.NotExistError(w, r, err)
			default:
				app.InternalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, postCtx, post)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *Application) BasicAuthMiddleware() func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				app.UnauthorizedBasicAuthError(w, r, fmt.Errorf("authorization header is missing"))
				return
			}

			parts := strings.Split(authHeader, " ")

			if len(parts) != 2 || parts[0] != "Basic" {
				app.UnauthorizedBasicAuthError(w, r, fmt.Errorf("authorization header is malformed"))
				return
			}
			decode, err := base64.StdEncoding.DecodeString(parts[1])
			if err != nil {
				app.UnauthorizedBasicAuthError(w, r, err)
				return
			}
			username := app.Config.Auth.BasicAuth.Username
			password := app.Config.Auth.BasicAuth.Password
			cred := strings.SplitN(string(decode), ":", 2)
			if len(cred) != 2 || cred[0] != username || cred[1] != password {
				app.UnauthorizedBasicAuthError(w, r, fmt.Errorf("invalid credentials"))
				return
			}

			next.ServeHTTP(w, r)

		})
	}

}

func (app *Application) CheckPostAuthorization(requiredRole string, next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := getUserFromContext(r)
		post := getPostFromContext(r)

		if user.ID == post.UserID {
			next.ServeHTTP(w, r)
			return
		}

		allowed, err := app.checkRolePrecedent(r.Context(), user, requiredRole)
		if err != nil {
			app.InternalServerError(w, r, err)
			return
		}

		if !allowed {
			app.ForbiddenError(w, r, err)
			return
		}

		next.ServeHTTP(w, r)

	})

}

func (app *Application) checkRolePrecedent(ctx context.Context, user *store.User, roleName string) (bool, error) {

	role, err := app.Store.Role.GetRoleByName(ctx, roleName)
	if err != nil {
		return false, err
	}
	return user.Role.Level >= role.Level, nil
}
