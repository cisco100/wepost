package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"

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
	plainToken := uuid.New().String()
	hash := sha256.Sum256([]byte(plainToken))
	hashToken := hex.EncodeToString(hash[:])
	type UserPayload struct {
		Username string `json:"username" validate:"required,max=100"`
		Email    string `json:"email" validate:"required,email,max=255"`
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

	err := app.Store.User.CreateAndInvite(ctx, user, hashToken, app.Config.Mail.InviteExpiry)
	if err != nil {
		switch err {
		case store.ErrDuplicateEmail:
			app.BadRequestError(w, r, err)

		case store.ErrDuplicateUsername:
			app.BadRequestError(w, r, err)

		default:
			app.InternalServerError(w, r, err)

		}
		return
	}

	type UserToken struct {
		User  *store.User `json:"user"`
		Token string      `json:"token"`
	}

	var userToken = UserToken{
		User:  user,
		Token: plainToken,
	}

	if err := JSONResponse(w, http.StatusCreated, userToken); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}

func (app *Application) ActivateUser(w http.ResponseWriter, r *http.Request) {

	token := chi.URLParam(r, "token")

	ctx := r.Context()

	err := app.Store.User.ActivateAccount(ctx, string(token), time.Now())

	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.NotExistError(w, r, err)
			return
		default:
			app.InternalServerError(w, r, err)
			return
		}
	}

	// type Message struct {
	// 	Msg string `json:"msg"`
	// }

	if err := JSONResponse(w, http.StatusNoContent, ""); err != nil {
		app.InternalServerError(w, r, err)
		return
	}

}
