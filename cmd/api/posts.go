package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/cisco100/wepost/internal/store"
)

func (app *Application) CreatePost(w http.ResponseWriter, r *http.Request) {

	type PostPayload struct {
		Title   string   `json:"title" validate:"required,min=3,max=100"`
		Content string   `json:"content" validate:"required,min=10"`
		Tags    []string `json:"tags" validate:"required,min=1,dive,required,min=1"`
	}

	var payload PostPayload
	if err := ReadJSON(w, r, &payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	post := &store.Post{
		ID:      uuid.New().String(),
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		UserID:  "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
	}

	ctx := r.Context()
	if err := app.Store.Post.Create(ctx, post); err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	if err := JSONResponse(w, http.StatusCreated, post); err != nil {
		app.InternalServerError(w, r, err)
		return
	}

}
func (app *Application) GetPostById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")

	ctx := r.Context()
	post, err := app.Store.Post.GetPostById(ctx, string(idParam))

	if err != nil {
		app.NotExistError(w, r, err)
	}

	comments, err := app.Store.Comment.GetPostWithComment(ctx, string(idParam))
	if err != nil {
		app.InternalServerError(w, r, err)
	}
	post.Comment = comments
	if err := JSONResponse(w, http.StatusFound, post); err != nil {
		app.InternalServerError(w, r, err)
	}
}

func (app *Application) GetAllPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts, err := app.Store.Post.AllPost(ctx)
	if err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	if err := JSONResponse(w, http.StatusOK, posts); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}

func (app *Application) DeletePost(w http.ResponseWriter, r *http.Request) {
	idParams := chi.URLParam(r, "postID")
	ctx := r.Context()

	if err := app.Store.Post.DeletePost(ctx, string(idParams)); err != nil {
		app.NotExistError(w, r, err)
	}

	type DeleteMsg struct {
		PostID string `json:"post_id"`
		Status int    `json:""status`
	}
	var msg DeleteMsg
	msg = DeleteMsg{PostID: string(idParams), Status: http.StatusNoContent}

	if err := JSONResponse(w, http.StatusOK, msg); err != nil {
		app.InternalServerError(w, r, err)
	}
}
