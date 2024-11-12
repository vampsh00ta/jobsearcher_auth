package http

import (
	isrvc "jobsearcher_user/internal/app/service"

	"github.com/gofiber/fiber/v2"
)

type authRoute struct {
	authSrvc isrvc.Auth
}

func newAuthRoute(app fiber.Router, authSrvc isrvc.Auth) {
	r := &authRoute{
		authSrvc: authSrvc,
	}

	app.Get("/accept", r.acceptToken)

}

type acceptTokenRequest struct {
	Code string `json:"code"`
}
type saveFilterResponse struct {
	Status string `json:"status"`
}

func (u authRoute) acceptToken(ctx *fiber.Ctx) error {
	var req acceptTokenRequest

	if err := ctx.QueryParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responseError{
			Error: err,
		})
	}
	//slugsToKeyword := func(slugs []string) []entity.Keyword {
	//	res := make([]entity.Keyword, len(slugs))
	//	for i := 0; i < len(slugs); i++ {
	//		res[i].Slug = slugs[i]
	//	}
	//	return res
	//}
	//filter := entity.Filter{
	//	UserTgID: req.UserTgID,
	//	City:     req.City,
	//	Company: entity.Company{
	//		Slug: req.Company,
	//	},
	//	Experience: entity.Experience{Slug: req.Experience},
	//	Keywords:   slugsToKeyword(req.KeywordsSlugs),
	//}
	//if err := u.userSrvc.SaveFilter(ctx.Context(), filter); err != nil {
	//	return ctx.Status(fiber.StatusInternalServerError).JSON(responseError{
	//		Error: err,
	//	})
	//}
	//res := saveFilterResponse{"ok"}
	return ctx.Status(fiber.StatusCreated).JSON(nil)
}

//
//
//
//type updateNotifyRequest struct {
//	TgID   int  `json:"tg_id"`
//	Option bool `json:"option"`
//}
//type updateNotifyResponse struct {
//	Status string `json:"status"`
//}
//
//func (u userRoute) updateNotify(ctx *fiber.Ctx) error {
//	var req updateNotifyRequest
//
//	if err := ctx.BodyParser(&req); err != nil {
//		return ctx.Status(fiber.StatusInternalServerError).JSON(responseError{
//			Error: err,
//		})
//	}
//
//	if err := u.userSrvc.UpdateNotify(ctx.Context(), req.TgID, req.Option); err != nil {
//		return ctx.Status(fiber.StatusInternalServerError).JSON(responseError{
//			Error: err,
//		})
//	}
//	res := updateNotifyResponse{"ok"}
//	return ctx.Status(fiber.StatusCreated).JSON(res)
//}
