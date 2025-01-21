package http

import (
	"github.com/go-playground/validator"
	srvc "jobsearcher_auth/internal/app/service"

	"github.com/gofiber/fiber/v2"
)

type Services struct {
	Auth srvc.Auth
	Link srvc.Link
}

func New(app *fiber.App, services Services) {
	prefix := "/api/v1"
	router := app.Group(prefix)
	validate := validator.New()
	newAuthRoute(router, services.Auth, services.Link, validate)
}
