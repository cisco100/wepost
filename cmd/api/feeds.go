package main

import "net/http"

func (app *Application) GetUserFeed(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	feeds, err := app.Store.Post.GetUserFeed(ctx, userID)
	if err != nil {
		app.InternalServerError(w, r, err)
		return
	}
	if err := JSONResponse(w, http.StatusOK, feeds); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}
