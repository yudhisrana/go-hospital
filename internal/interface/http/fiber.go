package http

import (
	"context"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/yudhisrana/go-hospital/internal/infra/config"
)

type Server struct {
	app *fiber.App
}

func NewServer(cfg config.AppConfig) *Server {
	app := fiber.New(fiber.Config{
		AppName:       cfg.AppName,
		ServerHeader:  cfg.AppName,
		ReadTimeout:   cfg.AppReadTimeout,
		WriteTimeout:  cfg.AppWriteTimeout,
		IdleTimeout:   cfg.AppIdleTimeout,
		StrictRouting: true,
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: splitAndTrim(cfg.AllowOrigins),
		AllowMethods: splitAndTrim(cfg.AllowMethods),
		AllowHeaders: splitAndTrim(cfg.AllowHeaders),
	}))

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

func splitAndTrim(value string) []string {
	if value == "" {
		return nil
	}

	parts := strings.Split(value, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, part)
		}
	}

	return result
}
