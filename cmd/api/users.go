package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/cisco100/wepost/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type userKey string

type UserToken struct {
	User  *store.User `json:"user"`
	Token string      `json:"token"`
}

type UserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"min=3,max=72"`
}

type TokenAuthPayload struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"min=3,max=72"`
}

const userCtx userKey = "user"

func getUserFromContext(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtx).(*store.User)
	return user
}

// @Summary		Get user by ID
// @Description	Get user by ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			userID	path		string	true	"User ID"
// @Success		200		{object}	store.User
// @Failure		404		{object}	error
// @Failure		500		{object}	error
// @Router			/users/getuser/{userID} [get]
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

// @Summary		Register a new user
// @Description	Register a new user
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			user	body		UserPayload	true	"User Payload"
// @Success		201		{object}	UserToken
// @Failure		400		{object}	error
// @Failure		500		{object}	error
// @Router			/register/user [post]
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

	// activationURL := fmt.Sprintf("%s/confirm/%s", app.Config.FrontendURL, plainToken)
	// isProdEnv := app.Config.Environment == "Production"
	// data := struct {
	// 	Username      string
	// 	ActivationURL string
	// }{
	// 	Username:      user.Username,
	// 	ActivationURL: activationURL,
	// }

	// if err := app.Mailer.Send(mailer.UserInvitesTEmplate, user.Username, user.Email, data, !isProdEnv); err != nil {
	// 	app.Logger.Errorw("error sending user activation email ::", "error", err)

	// 	if err := app.Store.User.DeleteUser(ctx, user.ID); err != nil {
	// 		app.Logger.Errorw("error deleting user  ::", "error", err)

	// 	}
	// 	app.InternalServerError(w, r, err)
	// 	return
	// }

	if err := JSONResponse(w, http.StatusCreated, userToken); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}

// @Summary		Activate a user account
// @Description	Activate a user account
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			token	path	string	true	"Activation Token"
// @Success		204
// @Failure		404	{object}	error
// @Failure		500	{object}	error
// @Router			/users/user/account/activate/{token} [post]
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

	if err := JSONResponse(w, http.StatusNoContent, ""); err != nil {
		app.InternalServerError(w, r, err)
		return
	}

}

// @Summary		Authenticate user and get token
// @Description	Authenticate user and get token
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			user	body		TokenAuthPayload	true	"User Credentials"
// @Success		201		{string}	string				"JWT Token"
// @Failure		400		{object}	error
// @Failure		401		{object}	error
// @Failure		500		{object}	error
// @Router			/auth/token-auth [post]
func (app *Application) TokenAuth(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	type TokenAuthPayload struct {
		Email    string `json:"email" validate:"required,email,max=255"`
		Password string `json:"password" validate:"min=3,max=72"`
	}
	var payload TokenAuthPayload
	if err := ReadJSON(w, r, &payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	user, err := app.Store.User.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.UnauthorizedError(w, r, err) //Didnt use` NotFound because want to avoid enummeration attack`
			return
		default:
			app.InternalServerError(w, r, err)
			return
		}
	}
	claims := jwt.MapClaims{
		"subs": user.ID,
		"exp":  time.Now().Add(app.Config.Auth.TokenAuth.Expiry).Unix(),
		"iat":  time.Now().Unix(),
		"nbf":  time.Now().Unix(),
		"iss":  app.Config.Auth.TokenAuth.Issue,
		"aud":  app.Config.Auth.TokenAuth.Audience,
	}
	token, err := app.Authenticator.GenerateToken(claims)
	if err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	if err := JSONResponse(w, http.StatusCreated, token); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}
