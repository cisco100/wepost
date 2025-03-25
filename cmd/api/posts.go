package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/cisco100/wepost/internal/store"
)

type PostPayload struct {
	Title   string   `json:"title" validate:"required,min=3,max=100"`
	Content string   `json:"content" validate:"required,min=10"`
	Tags    []string `json:"tags"`
}

type PostUpdatePayload struct {
	Title   *string  `json:"title" validate:"omitempty,min=3,max=100"`
	Content *string  `json:"content" validate:"omitempty,min=10"`
	Tags    []string `json:"tags" validate:"omitempty"`
}

// @Summary		Creates a post
// @Description	Creates a post
// @Tags			posts
// @Accept			json
// @Produce		json
// @Param			payload	body		PostPayload	true	"Post payload"
// @Success		201		{object}	store.Post
// @Failure		400		{object}	error
// @Failure		401		{object}	error
// @Failure		500		{object}	error
// @Security		ApiKeyAuth
// @Router			/posts/create-post [post]
func (app *Application) CreatePost(w http.ResponseWriter, r *http.Request) {

	type PostPayload struct {
		Title   string   `json:"title" validate:"required,min=3,max=100"`
		Content string   `json:"content" validate:"required,min=10"`
		Tags    []string `json:"tags"`
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
	user := getUserFromContext(r)

	post := &store.Post{
		ID:      uuid.New().String(),
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		UserID:  user.ID,
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

// @Summary		Fetches a post
// @Description	Fetches a post by ID
// @Tags			posts
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Post ID"
// @Success		200	{object}	store.Post
// @Failure		404	{object}	error
// @Failure		500	{object}	error
// @Security		ApiKeyAuth
// @Router			/posts/getpost/{postID} [get]
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

// @Summary		Fetches all posts
// @Description	Fetches all posts
// @Tags			posts
// @Accept			json
// @Produce		json
// @Success		200	{object}	store.Post
// @Failure		404	{object}	error
// @Failure		500	{object}	error
// @Security		ApiKeyAuth
// @Router			/posts/all [get]
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

// @Summary		Deletes a post
// @Description	Delete a post by ID
// @Tags			posts
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Post ID"
// @Success		204	{object}	string
// @Failure		404	{object}	error
// @Failure		500	{object}	error
// @Security		ApiKeyAuth
// @Router			/posts/delete/{postID} [delete]
func (app *Application) DeletePost(w http.ResponseWriter, r *http.Request) {
	idParams := chi.URLParam(r, "postID")
	ctx := r.Context()

	if err := app.Store.Post.DeletePost(ctx, string(idParams)); err != nil {
		app.NotExistError(w, r, err)
	}

	type DeleteMsg struct {
		PostID string `json:"post_id"`
		Status int    `json:"status"`
	}
	msg := DeleteMsg{PostID: string(idParams), Status: http.StatusNoContent}

	if err := JSONResponse(w, http.StatusOK, msg); err != nil {
		app.InternalServerError(w, r, err)
	}
}

// @Summary		Updates a post
// @Description	Updates a post by ID
// @Tags			posts
// @Accept			json
// @Produce		json
// @Param			id		path		string				true	"Post ID"
// @Param			payload	body		PostUpdatePayload	true	"Post update payload"
// @Success		200		{object}	store.Post
// @Failure		400		{object}	error
// @Failure		404		{object}	error
// @Failure		500		{object}	error
// @Security		ApiKeyAuth
// @Router			/posts/post/update/{postID} [put]
func (app *Application) UpdatePost(w http.ResponseWriter, r *http.Request) {
	idParams := chi.URLParam(r, "postID")
	ctx := r.Context()
	type PostUpdatePayload struct {
		Title   *string  `json:"title" validate:"omitempty,min=3,max=100"`
		Content *string  `json:"content" validate:"omitempty,min=10"`
		Tags    []string `json:"tags" validate:"omitempty"`
	}

	var payload PostUpdatePayload

	if err := ReadJSON(w, r, &payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	existingPost, err := app.Store.Post.GetPostById(ctx, string(idParams))
	if err != nil {
		app.NotExistError(w, r, err)
		return

	}
	if payload.Title != nil {
		existingPost.Title = *payload.Title

	}
	if payload.Content != nil {
		existingPost.Content = *payload.Content

	}
	if payload.Tags != nil {
		existingPost.Tags = payload.Tags

	}

	if err := app.Store.Post.UpdatePost(ctx, existingPost); err != nil {
		app.InternalServerError(w, r, err)
	}

	if err := JSONResponse(w, http.StatusOK, existingPost); err != nil {
		app.InternalServerError(w, r, err)
	}

}
