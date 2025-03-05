package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// GetUser godoc
//
//	@Summary		Fetches a user profile
//	@Description	Fetches a user profile by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{object}	store.User
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Security		ApiKeyAuth
//	@Router			/users/getuser/{id} [get]
func (app *Application) GetUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "userID")
	ctx := r.Context()
	user, err := app.Store.User.GetUserById(ctx, string(idParam))

	if err != nil {
		app.NotExistError(w, r, err)
	}

	if err := JSONResponse(w, http.StatusFound, user); err != nil {
		app.InternalServerError(w, r, err)
	}
}
