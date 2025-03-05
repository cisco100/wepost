package main

import (
	"net/http"

	"github.com/cisco100/wepost/internal/store"
	"github.com/go-chi/chi/v5"
)

type FollowPayload struct {
	FollowerID string `json:"follower_id"`
}

// FollowUser godoc
//
//	@Summary		Follows a user
//	@Description	Follows a user by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userID	path		string	true	"User ID"
//	@Success		204		{string}	string	"User followed"
//	@Failure		400		{object}	error	"User payload missing"
//	@Failure		404		{object}	error	"User not found"
//	@Security		ApiKeyAuth
//	@Router			/users/getuser/{userID}/follow [put]
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

// UnfollowUser gdoc
//
//	@Summary		Unfollow a user
//	@Description	Unfollow a user by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userID	path		int		true	"User ID"
//	@Success		204		{string}	string	"User unfollowed"
//	@Failure		400		{object}	error	"User payload missing"
//	@Failure		404		{object}	error	"User not found"
//	@Security		ApiKeyAuth
//	@Router			/users/{userID}/unfollow [put]

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
