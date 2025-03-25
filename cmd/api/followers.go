package main

import (
	"net/http"

	"github.com/cisco100/wepost/internal/store"
	"github.com/go-chi/chi/v5"
)

type FollowPayload struct {
	FollowerID string `json:"follower_id"`
}

// @Summary		Follow a user
// @Description	Follow a user by ID
// @Tags			followers
// @Accept			json
// @Produce		json
// @Param			userID	path		string			true	"User ID"
// @Param			payload	body		FollowPayload	true	"Follow payload"
// @Success		200		{object}	nil
// @Failure		400		{object}	error
// @Failure		404		{object}	error
// @Failure		409		{object}	error
// @Failure		500		{object}	error
// @Security		ApiKeyAuth
// @Router			/users/getuser/{userID}/follow [post]
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

// @Summary		Unfollow a user
// @Description	Unfollow a user by ID
// @Tags			followers
// @Accept			json
// @Produce		json
// @Param			userID	path		string			true	"User ID"
// @Param			payload	body		FollowPayload	true	"Unfollow payload"
// @Success		204		{object}	nil
// @Failure		400		{object}	error
// @Failure		404		{object}	error
// @Failure		409		{object}	error
// @Failure		500		{object}	error
// @Security		ApiKeyAuth
// @Router			/users/getuser/{userID}/unfollow [post]
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
