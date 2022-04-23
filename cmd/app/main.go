package main

import (
	"log"
	"os"

	"github.com/migmatore/college-site-backend/internal/app"
)

func main() {
	dsn := "host=127.0.0.1 user=postgres password= database=postgres"
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	app.Run(port, dsn)
}
