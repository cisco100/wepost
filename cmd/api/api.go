package main

import (
	"net/http"
	"os"
	"time"

	"github.com/cisco100/wepost/docs"
	"github.com/cisco100/wepost/internal/authenticator"
	"github.com/cisco100/wepost/internal/mailer"
	"github.com/cisco100/wepost/internal/store"
	"github.com/cisco100/wepost/internal/store/cache"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type Application struct {
	Config        AppConfig
	Store         store.Storage
	Logger        *zap.SugaredLogger
	Mailer        mailer.Client
	Authenticator authenticator.Authenticator
	Cache         cache.Storage
}

type AppConfig struct {
	Address     string
	Database    DbConfig
	Version     string
	Environment string
	APIURL      string
	Mail        MailConfig
	FrontendURL string
	Auth        AuthConfig
	Redis       RedisConfig
}

type RedisConfig struct {
	Addr     string
	Password string
	Database int
	Enabled  bool
}

type AuthConfig struct {
	BasicAuth BasicAuthConfig
	TokenAuth TokenAuthConfig
}

type TokenAuthConfig struct {
	Secret   string
	Audience string
	Issue    string
	Expiry   time.Duration
}

type BasicAuthConfig struct {
	Username string
	Password string
}

type MailConfig struct {
	SendGrid     SendgridConfig
	InviteExpiry time.Duration
}
type SendgridConfig struct {
	From   string
	ApiKey string
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

	app.Logger.Infow("Server started", "address", srv.Addr, "env", app.Config.Environment)

	return srv.ListenAndServe()
}
