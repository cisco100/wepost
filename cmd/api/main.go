package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/cisco100/wepost/internal/db"
	"github.com/cisco100/wepost/internal/mailer"
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

// package main

// import (
//     "fmt"
//     gomail "gopkg.in/mail.v2"
// )

// func main() {
//     // Create a new message
//     message := gomail.NewMessage()

//     // Set email headers
//     message.SetHeader("From", "youremail@email.com")
//     message.SetHeader("To", "recipient1@email.com")
//     message.SetHeader("Subject", "Hello from the Mailtrap team")

//     // Set email body
//     message.SetBody("text/plain", "This is the Test Body")

//     // Set up the SMTP dialer
//     dialer := gomail.NewDialer("live.smtp.mailtrap.io", 587, "api", "1a2b3c4d5e6f7g")

//     // Send the email
//     if err := dialer.DialAndSend(message); err != nil {
//         fmt.Println("Error:", err)
//         panic(err)
//     } else {
//         fmt.Println("Email sent successfully!")
//     }
// }

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
	api_key := os.Getenv("SENDGRID_API_KEY")
	from := os.Getenv("FROM_EMAIL")
	sndg := SendgridConfig{From: from, ApiKey: api_key}
	mail := MailConfig{SendGrid: sndg, InviteExpiry: time.Hour * 24} //  a day
	url := "http://localhost/4000"
	dbConfig := DbConfig{Addr: addr, MaxOpenConn: maxOpenCon, MaxIdleConn: maxIdleCon, MaxIdleTime: maxIdletime}
	ver := os.Getenv("VERSION")
	env := os.Getenv("ENVIRONMENT")
	basicAuth := BasicAuthConfig{Username: os.Getenv("AUTH_BASIC_USER"), Password: os.Getenv("AUTH_BASIC_PASSWORD")}
	auth := AuthConfig{BasicAuth: basicAuth}
	cfg := AppConfig{
		Address:     port,
		Database:    dbConfig,
		Version:     ver,
		Environment: env,
		APIURL:      api,
		Mail:        mail,
		FrontendURL: url,
		Auth:        auth,
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
	app := &Application{
		Config: cfg,
		Store:  store,
		Logger: logger,
		Mailer: mailer,
	}
	mux := app.Mount()
	logger.Fatal(app.Run(mux))
}
