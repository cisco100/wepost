package main

import (
	"net/http"

	"github.com/cisco100/wepost/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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

func (app *Application) RegisterUser(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	type UserPayload struct {
		Username string `json:"username" validate:"required,max=100"`
		Email    string `json:"email" validate:"required email,max=255"`
		Password string `json:"password" validate:"min=3,max=72"`
	}

	var payload UserPayload

	if err := ReadJSON(w, r, &payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}
	if err := Validate.Struct(payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	user := &store.User{
		ID:       uuid.New().String(),
		Username: payload.Username,
		Email:    payload.Email,
	}

	if err := user.Password.Set(payload.Password); err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	if err := app.Store.User.CreateAndInvite(ctx, user, uuid.New().String()); err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	if err := JSONResponse(w, http.StatusOK, nil); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}
