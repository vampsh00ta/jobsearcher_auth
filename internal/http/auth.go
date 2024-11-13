package http

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	isrvc "jobsearcher_user/internal/app/service"
	"jobsearcher_user/internal/entity"
)

type authRoute struct {
	authSrvc isrvc.Auth
	linkSrvc isrvc.Link
	validate *validator.Validate
}

func newAuthRoute(app fiber.Router, authSrvc isrvc.Auth, linkSrvc isrvc.Link, validate *validator.Validate) {
	r := &authRoute{
		authSrvc: authSrvc,
		linkSrvc: linkSrvc,
		validate: validate,
	}
	prefix := "/auth"
	app.Post(prefix+"/accept", r.acceptToken)
	app.Post(prefix+"/verify", r.verifyToken)
	app.Post(prefix+"/create_link", r.createLink)

}

type verifyTokenRequest struct {
	AccessToken string `json:"access_token" validate:"required"`
}
type verifyTokenResponse struct {
	Status bool  `json:"status"`
	Err    error `json:"error"`
}

func (u authRoute) verifyToken(ctx *fiber.Ctx) error {
	var req verifyTokenRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responseError{
			Error: err,
		})
	}
	//resp, err := http.Post("localhost:8081/api/v1/auth/accept",)
	//if err != nil {
	//	return nil, err
	//}
	////defer resp.Body.Close()
	//
	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	return nil, err
	//}
	if err := u.validate.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responseError{
			Error: err,
		})
	}
	ok, err := u.authSrvc.VerifyToken(ctx.Context(), req.AccessToken)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			verifyTokenResponse{
				Status: ok,
				Err:    err,
			})
	}

	return ctx.Status(fiber.StatusCreated).JSON(verifyTokenResponse{Status: ok})
}

type acceptTokenRequest struct {
	Hash string `json:"hash" validate:"required"`
}
type acceptTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (u authRoute) acceptToken(ctx *fiber.Ctx) error {
	var req acceptTokenRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responseError{
			Error: err,
		})
	}
	fmt.Println(req)
	if err := u.validate.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responseError{
			Error: err,
		})
	}
	access, err := u.linkSrvc.Claim(ctx.Context(), req.Hash)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(responseError{
			Error: err,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(acceptTokenResponse{AccessToken: access})
}

type createLinkRequest struct {
	ID        int    `json:"id" validate:"required"`
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name"`
	Username  string `json:"username"  validate:"required"`
	PhotoUrl  string `json:"photo_url"`
}
type createLinkResponse struct {
	Hash string `json:"hash"`
}

func (u authRoute) createLink(ctx *fiber.Ctx) error {
	var req createLinkRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responseError{
			Error: err,
		})
	}

	if err := u.validate.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	user := entity.User{
		req.ID,
		req.FirstName,
		req.LastName,
		req.Username,
		req.PhotoUrl,
	}
	access, err := u.authSrvc.CreateToken(ctx.Context(), user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	hash, err := u.linkSrvc.Create(ctx.Context(), req.Username, access)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(responseError{
			Error: err,
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(createLinkResponse{Hash: hash})
}
