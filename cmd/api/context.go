package main

import (
	"net/http"

	"github.com/cisco100/wepost/internal/store"
)

type postKey string
type userKey string

const postCtx postKey = "post"

const userCtx userKey = "user"

func getPostFromContext(r *http.Request) *store.Post {
	post, ok := r.Context().Value(postCtx).(*store.Post)
	if !ok {
		return nil
	}
	return post
}

func getUserFromContext(r *http.Request) *store.User {
	user, ok := r.Context().Value(userCtx).(*store.User)
	if !ok {
		return nil
	}
	return user
}
