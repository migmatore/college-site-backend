package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/migmatore/college-site-backend/internal/app"
)

func main() {
	_url := os.Getenv("DATABASE_URL")

	db, err := url.Parse(_url)
	if err != nil {
		panic(err)
	}

	p, _ := db.User.Password()

	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s", db.Hostname(), db.User.Username(), p, db.Path)
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	app.Run(port, dsn)
}
