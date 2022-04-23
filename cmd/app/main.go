package main

import (
	"github.com/migmatore/college-site-backend/internal/app"
)

func main() {
	dsn := "host=127.0.0.1 user=postgres password=KazakovVova45 database=postgres"

	app.Run(dsn)
}
