package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cisco100/wepost/docs"
	"github.com/cisco100/wepost/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Application struct {
	Config AppConfig
	Store  store.Storage
}

type AppConfig struct {
	Address     string
	Database    DbConfig
	Version     string
	Environment string
	APIURL      string
}

type DbConfig struct {
	Addr        string
	MaxOpenConn int
	MaxIdleConn int
	MaxIdleTime string
}

func (app *Application) Mount() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	RegisterRoutes(router, app)
	return router
}

func (app *Application) Run(mux http.Handler) error {
	docs.SwaggerInfo.Version = os.Getenv("VERSION")
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Host = app.Config.APIURL
	srv := &http.Server{
		Addr:         app.Config.Address,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Server started at port %s", srv.Addr)

	return srv.ListenAndServe()
}
