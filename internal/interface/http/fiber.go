package http

import (
	"context"

	"github.com/gofiber/fiber/v3"
)

type Server struct {
	app *fiber.App
}

func NewServer() *Server {
	app := fiber.New(fiber.Config{
		AppName: "Go-Hospital",
	})

	return &Server{app: app}
}

func (s *Server) App() *fiber.App {
	return s.app
}

func (s *Server) Start(port string) error {
	return s.app.Listen(":" + port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}
