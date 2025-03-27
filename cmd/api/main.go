package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/cisco100/wepost/internal/authenticator"
	"github.com/cisco100/wepost/internal/db"
	"github.com/cisco100/wepost/internal/mailer"
	"github.com/cisco100/wepost/internal/store"
	"github.com/cisco100/wepost/internal/store/cache"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

//	@title			WePost API
//	@version		1.0
//	@description	Social app for all.
//	@termsOfService	https://github.com/cisco100/wepost/blob/main/README.md

//	@contact.name	API Support
//	@contact.url	https://github.com/cisco100/wepost/issues/
//	@contact.email	web04501@gmail.com

//	@license.name	MIT
//	@license.url	https://github.com/cisco100/wepost/blob/main/LICENSE

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description

func main() {
	rootPath := "/home/cisco/pro/go/wepost"
	// Load the .env file from the root directory
	err := godotenv.Load(rootPath + "/.env")
	if err != nil {
		log.Fatalf("Error loading .env file from root: %v", err)
	}

	raddr := os.Getenv("REDIS_ADDR")
	rpwd := os.Getenv("REDIS_PASSWORD")
	rdbse, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	renb, _ := strconv.ParseBool(os.Getenv("REDIS_ENABLED"))

	port := os.Getenv("ADDR")
	addr := os.Getenv("DB_ADDR")
	maxOpenCon, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN"))
	maxIdleCon, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))
	maxIdletime := os.Getenv("DB_MAX_IDLE_TIME")
	api := os.Getenv("APIURL")
	api_key := os.Getenv("SENDGRID_API_KEY")
	from := os.Getenv("FROM_EMAIL")
	sndg := SendgridConfig{From: from, ApiKey: api_key}
	mail := MailConfig{SendGrid: sndg, InviteExpiry: time.Hour * 24} //  a day
	url := "http://localhost/4000"
	dbConfig := DbConfig{Addr: addr, MaxOpenConn: maxOpenCon, MaxIdleConn: maxIdleCon, MaxIdleTime: maxIdletime}
	ver := os.Getenv("VERSION")
	env := os.Getenv("ENVIRONMENT")
	basicAuth := BasicAuthConfig{Username: os.Getenv("AUTH_BASIC_USER"), Password: os.Getenv("AUTH_BASIC_PASSWORD")}
	issuer := "wepost"
	audience := "wepost"

	token := TokenAuthConfig{Secret: os.Getenv("AUTH_TOKEN_SECRET"), Audience: audience, Issue: issuer, Expiry: time.Hour * 24 * 2}
	auth := AuthConfig{BasicAuth: basicAuth, TokenAuth: token}
	red := RedisConfig{Addr: raddr, Password: rpwd, Database: rdbse, Enabled: renb}
	cfg := AppConfig{
		Address:     port,
		Database:    dbConfig,
		Version:     ver,
		Environment: env,
		APIURL:      api,
		Mail:        mail,
		FrontendURL: url,
		Auth:        auth,
		Redis:       red,
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()
	db, err := db.NewConnection(
		cfg.Database.Addr,
		cfg.Database.MaxOpenConn,
		cfg.Database.MaxIdleConn,
		cfg.Database.MaxIdleTime,
	)
	if err != nil {
		logger.Panic(err)
	}
	defer db.Close()
	logger.Info("Database connection pool successfuly established")
	mailer := mailer.NewSendGridMailer(cfg.Mail.SendGrid.From, cfg.Mail.SendGrid.ApiKey)
	store := store.NewStorage(db)

	var rdb *redis.Client
	if cfg.Redis.Enabled {
		rdb = cache.NewRedisConnection(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.Database)
		logger.Info("Redis connection pool  successfuly establshed")
	}
	cache := cache.NewRedisStorage(rdb)

	jwtAuthenticator := authenticator.NewJWTAuthenticator(
		cfg.Auth.TokenAuth.Secret,
		cfg.Auth.TokenAuth.Audience,
		cfg.Auth.TokenAuth.Issue,
		cfg.Auth.TokenAuth.Expiry,
	)
	app := &Application{
		Config:        cfg,
		Store:         store,
		Logger:        logger,
		Mailer:        mailer,
		Authenticator: jwtAuthenticator,
		Cache:         cache,
	}
	mux := app.Mount()
	logger.Fatal(app.Run(mux))
}
