package main

import "net/http"

func (app *Application) GetUserFeed(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := "957ef048-1d04-4b0a-9d0c-7e690a0a652a"
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
