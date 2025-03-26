package main

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

func RegisterRoutes(router chi.Router, app *Application) {
	router.Route("/v1/", func(r chi.Router) {
		URL := fmt.Sprintf("%s/swagger/doc.json", app.Config.Address)

		//=============INFO ROUTES===========//
		r.With(app.BasicAuthMiddleware()).Get("/info", app.Meta)
		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(URL),
		))
		//=============INFO ROUTES===========//

		//=============PUBLIC ROUTES===========//
		r.Post("/register/user", app.RegisterUser)
		r.Put("/users/user/account/activate/{token}", app.ActivateUser)
		r.Post("/auth/token-auth", app.TokenAuth)
		//=============PUBLIC ROUTES===========//

		//=============AUTHENTICATED ROUTES===========//
		r.Group(func(r chi.Router) {
			r.Use(app.PostsContextMiddleware)
			r.Use(app.AuthTokenMiddleware())

			//=============POST ROUTES===========//
			r.Post("/posts/create-post", app.CreatePost)
			r.Post("/posts/comments/create-comment", app.CreateComment)
			r.Patch("/posts/post/update/{postID}", app.CheckPostAuthorization("moderator", app.UpdatePost))
			r.Delete("/posts/delete/{postID}", app.CheckPostAuthorization("admin", app.DeletePost))
			r.Get("/posts/getpost/{postID}", app.GetPostById)
			r.Get("/posts/all", app.GetAllPost)
			//=============POST ROUTES===========//

			//=============USER ROUTES===========//
			r.Get("/users/getuser/{userID}", app.GetUserById)
			r.Get("/users/feeds", app.GetUserFeed)
			r.Put("/users/getuser/{userID}/follow", app.FollowUser)
			r.Put("/users/getuser/{userID}/unfollow", app.UnFollowUser)
			//=============USER ROUTES===========//
		})
		//=============AUTHENTICATED ROUTES===========//
	})
}
