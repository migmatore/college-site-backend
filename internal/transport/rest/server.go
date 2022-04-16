package rest

import "github.com/gofiber/fiber/v2"

type Server struct {
	app  *fiber.App
	port string
}

func NewServer(port string, app *fiber.App) *Server {
	return &Server{
		app:  app,
		port: port,
	}
}

func (app *Server) Listen() {
	app.app.Listen(app.port)
}
