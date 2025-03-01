package main

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router chi.Router, app *Application) {
	router.Route("/v1/", func(r chi.Router) {
		//=============POST ROUTES===========//
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

		//=============POST ROUTES===========//

		//=============USER ROUTES===========//
		//Get URL
		r.Get("/users/getuser/{userID}", app.GetUserById)

		//Put URL
		r.Put("/users/getuser/{userID}/follow", app.FollowUser)
		r.Put("/users/getuser/{userID}/unfollow", app.UnFollowUser)

		//Post URL

		//Update URL

		// Delete URL

		//=============USER ROUTES===========//
	})

}
