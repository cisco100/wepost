package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

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
