package main

import (
	"log"
	"os"
	"strconv"

	"github.com/cisco100/wepost/internal/db"
	"github.com/cisco100/wepost/internal/store"
	"github.com/joho/godotenv"
)

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
	dbConfig := DbConfig{Addr: addr, MaxOpenConn: maxOpenCon, MaxIdleConn: maxIdleCon, MaxIdleTime: maxIdletime}
	ver := os.Getenv("VERSION")
	env := os.Getenv("ENVIRONMENT")
	cfg := AppConfig{
		Address:     port,
		Database:    dbConfig,
		Version:     ver,
		Environment: env,
	}

	db, err := db.NewConnection(
		cfg.Database.Addr,
		cfg.Database.MaxOpenConn,
		cfg.Database.MaxIdleConn,
		cfg.Database.MaxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	log.Println("Database connection pool successfuly established")

	store := store.NewStorage(db)
	app := &Application{
		Config: cfg,
		Store:  store,
	}
	mux := app.Mount()
	log.Fatal(app.Run(mux))
}
