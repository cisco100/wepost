package main

import (
	"log"
	"os"
	"strconv"

	"github.com/cisco100/wepost/internal/db"
	"github.com/cisco100/wepost/internal/store"
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
	port := os.Getenv("ADDR")
	addr := os.Getenv("DB_ADDR")
	maxOpenCon, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN"))
	maxIdleCon, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))
	maxIdletime := os.Getenv("DB_MAX_IDLE_TIME")
	api := os.Getenv("APIURL")
	dbConfig := DbConfig{Addr: addr, MaxOpenConn: maxOpenCon, MaxIdleConn: maxIdleCon, MaxIdleTime: maxIdletime}
	ver := os.Getenv("VERSION")
	env := os.Getenv("ENVIRONMENT")
	cfg := AppConfig{
		Address:     port,
		Database:    dbConfig,
		Version:     ver,
		Environment: env,
		APIURL:      api,
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

	store := store.NewStorage(db)
	app := &Application{
		Config: cfg,
		Store:  store,
		Logger: logger,
	}
	mux := app.Mount()
	logger.Fatal(app.Run(mux))
}
