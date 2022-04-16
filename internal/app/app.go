package app

import (
	"log"

	"github.com/migmatore/college-site-backend/internal/service"
	"github.com/migmatore/college-site-backend/internal/storage"
	"github.com/migmatore/college-site-backend/internal/storage/psql"
	"github.com/migmatore/college-site-backend/internal/transport/rest"
	"github.com/migmatore/college-site-backend/internal/transport/rest/handler"
)

func Run(dsn string) {
	db, err := psql.NewPostgress(dsn)
	if err != nil {
		log.Fatalf("failed to initialize db connetcion: %s", err.Error())
	}

	storages := storage.New(db)

	services := service.New(service.Deps{
		GroupStorage: storages.Group,
	})

	restHandlers := handler.New(handler.Deps{
		GroupService: services.Group,
	})

	app := restHandlers.Init()

	srv := rest.NewServer(":3000", app)
	srv.Listen()
}
