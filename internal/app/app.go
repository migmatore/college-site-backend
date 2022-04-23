package app

import (
	"log"

	"github.com/migmatore/college-site-backend/internal/service"
	"github.com/migmatore/college-site-backend/internal/storage"
	"github.com/migmatore/college-site-backend/internal/storage/psql"
	"github.com/migmatore/college-site-backend/internal/transport/rest"
	"github.com/migmatore/college-site-backend/internal/transport/rest/handler"
)

func Run(port string, dsn string) {
	db, err := psql.NewPostgress(dsn)
	if err != nil {
		log.Fatalf("failed to initialize db connetcion: %s", err.Error())
	}

	storages := storage.New(db)

	services := service.New(service.Deps{
		GroupStorage:   storages.Group,
		OfficeStorage:  storages.Office,
		SubjectStorage: storages.Subject,
		TeacherStorage: storages.Teacher,
		WeekdayStorage: storages.Weekday,
		SheduleStorage: storages.Shedule,
	})

	restHandlers := handler.New(handler.Deps{
		GroupService:   services.Group,
		OfficeService:  services.Office,
		SubjectService: services.Subject,
		TeacherService: services.Teacher,
		WeekdayService: services.Weekday,
		SheduleService: services.Shedule,
	})

	app := restHandlers.Init()

	srv := rest.NewServer("0.0.0.0:"+port, app)
	srv.Listen()
}
