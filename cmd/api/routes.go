package main

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router chi.Router, app *Application) {
	router.Route("/v1/", func(r chi.Router) {
		//Get URL
		r.Get("/info", app.Meta)
		r.Get("/posts/getpost/{postID}", app.GetPostById)
		r.Get("/posts/all", app.GetAllPost)

		//Post URL
		r.Post("/posts/create-post", app.CreatePost)
		r.Post("/posts/comments/create-comment", app.CreateComment)

		//Update URL
		r.Patch("/posts/post/update/{postID}", app.UpdatePost)

		// Delete URL
		r.Delete("/posts/delete/{postID}", app.DeletePost)
	})

}
