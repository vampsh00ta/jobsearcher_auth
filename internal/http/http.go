package http

import (
	srvc "jobsearcher_user/internal/app/service"

	"github.com/gofiber/fiber/v2"
)

type Services struct {
	Auth srvc.Auth
}

func New(app *fiber.App, services Services) {
	prefix := "/api/v1/auth"
	router := app.Group(prefix)

	newAuthRoute(router, services.Auth)
}
