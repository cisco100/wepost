package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

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
