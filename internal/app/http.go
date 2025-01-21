package app

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
	"jobsearcher_auth/config"
)

func newHTTP(cfg *config.Config, logger *zap.Logger) *fiber.App {
	app := fiber.New(fiber.Config{AppName: cfg.HTTP.Name})
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))
	// app.User(log)
	app.Use(compress.New())

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins: "*",
		//AllowCredentials: true,
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	return app
}
