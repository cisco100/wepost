package main

import (
	"net/http"

	"github.com/cisco100/wepost/internal/store"
)

func (app *Application) GetUserFeed(w http.ResponseWriter, r *http.Request) {

	feedPaginator := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	feedPaginator, err := feedPaginator.Parser(r)
	if err != nil {
		app.BadRequestError(w, r, err)
	}

	if err := Validate.Struct(feedPaginator); err != nil {
		app.BadRequestError(w, r, err)
	}

	ctx := r.Context()
	userID := "20de04a8-e55a-45c9-9462-743a176b02f1"
	feeds, err := app.Store.Post.GetUserFeed(ctx, userID, feedPaginator)
	if err != nil {
		app.InternalServerError(w, r, err)
		return
	}
	if err := JSONResponse(w, http.StatusOK, feeds); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}
