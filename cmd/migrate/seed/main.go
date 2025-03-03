package main

import (
	"log"
	"os"

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
	addr := os.Getenv("DB_ADDR")
	conn, err := db.NewConnection(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.SeedDB(store)
}
