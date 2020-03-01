package main

import (
	"context"
	"google-service/router"
	"os"
	"log"
	"google-service/driver"
)

// @title MIPT Service API
// @version 0.1
// @description This is MIPT microservice
// @termsOfService http://swagger.io/terms/

// @contact.name Digitory
// @contact.url http://digitory.dev
// @contact.email support@digitory.dev

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath /api/v1
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// CRON Shedule
// 	wg := &sync.WaitGroup{}
// 	wg.Add(1)
// 	c := cron.New()
//
// 	c.AddFunc("@every 15ь", handler.LDCHandler().SendDataToYCBM)
// 	c.Start()

    // ENV variables for database connection
	ctx = context.WithValue(ctx, driver.HostKey, os.Getenv("MONGO_HOST"))
	ctx = context.WithValue(ctx, driver.UsernameKey, os.Getenv("MONGO_USERNAME"))
	ctx = context.WithValue(ctx, driver.PasswordKey, os.Getenv("MONGO_PASSWORD"))
	ctx = context.WithValue(ctx, driver.DatabaseKey, os.Getenv("MONGO_DATABASE"))

	db, err := driver.Connect(ctx)
	if err != nil {
		log.Fatalf("Database configuration failed: %v", err)
	}

	r := router.SetupRouter(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

// 	wg.Wait()
}
