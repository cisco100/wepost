package main

import (
	"net/http"

	"github.com/cisco100/wepost/internal/store"
	"github.com/go-chi/chi/v5"
)

type FollowPayload struct {
	FollowerID string `json:"follower_id"`
}

func (app *Application) FollowUser(w http.ResponseWriter, r *http.Request) {
	var payload FollowPayload
	idParam := chi.URLParam(r, "userID")
	ctx := r.Context()
	if err := ReadJSON(w, r, &payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	user, err := app.Store.User.GetUserById(ctx, idParam)
	if err != nil {
		app.NotExistError(w, r, err)
		return
	}

	if err := app.Store.Follower.Follow(ctx, user.ID, payload.FollowerID); err != nil {
		switch err {
		case store.ErrConflict:
			app.ConflictError(w, r, err)
			return
		default:
			app.InternalServerError(w, r, err)
			return
		}

	}

	if err := JSONResponse(w, http.StatusOK, nil); err != nil {
		app.InternalServerError(w, r, err)
		return
	}

}

func (app *Application) UnFollowUser(w http.ResponseWriter, r *http.Request) {
	var payload FollowPayload
	idParam := chi.URLParam(r, "userID")
	ctx := r.Context()
	if err := ReadJSON(w, r, &payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	user, err := app.Store.User.GetUserById(ctx, idParam)
	if err != nil {
		app.NotExistError(w, r, err)
		return
	}

	if err := app.Store.Follower.Unfollow(ctx, user.ID, payload.FollowerID); err != nil {
		switch err {
		case store.ErrConflict:
			app.ConflictError(w, r, err)
			return
		default:
			app.InternalServerError(w, r, err)
			return
		}

	}

	if err := JSONResponse(w, http.StatusNoContent, nil); err != nil {
		app.InternalServerError(w, r, err)
	}

}
