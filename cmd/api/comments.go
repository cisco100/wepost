package main

import (
	"net/http"

	"github.com/cisco100/wepost/internal/store"
	"github.com/google/uuid"
)

func (app *Application) CreateComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type CommentPayload struct {
		Comment string `json:"comment" validate:"required,min=10"`
	}
	var payload CommentPayload
	if err := ReadJSON(w, r, &payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	comment := &store.Comment{
		ID:      uuid.New().String(),
		PostID:  "49cf699e-be1f-4734-80d1-96ef941d93af",
		UserID:  "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		Comment: payload.Comment,
	}

	if err := app.Store.Comment.Create(ctx, comment); err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	if err := JSONResponse(w, http.StatusCreated, comment); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}
