package auth

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shelllbyyyyy/belajar-api-go/internal/api"
	"github.com/shelllbyyyyy/belajar-api-go/internal/environtment"
	"github.com/shelllbyyyyy/belajar-api-go/internal/exception"
)

type authHandler struct {
    auth AuthUseCase
    user UserUseCase
}

func newAuthHandler(au AuthUseCase, uu UserUseCase) *authHandler {
    return &authHandler{
		auth: au,
		user: uu,
	}
}

func (h authHandler) Register(ctx *fiber.Ctx) error {
    var req = registerUserSchema{}

	if err := ctx.BodyParser(&req); err != nil {
		myErr := exception.ErrorBadRequest
		return api.NewResponse(
			api.WithMessage(err.Error()),
			api.WithError(myErr),
			api.WithHttpCode(http.StatusBadRequest),
			api.WithMessage("register fail"),
		).Send(ctx)
	}

	_, err := h.user.FindUserByEmail(ctx.UserContext(), req.Email)
	if err == nil {
		return api.NewResponse(
			api.WithMessage("Email has already been registered"),
			api.WithError(exception.ErrorEmailAlreadyUsed),
			api.WithHttpCode(http.StatusConflict),
		).Send(ctx)
	}

     result, err := h.auth.CreateUser(ctx.UserContext(), req) 
	 if err != nil {
		myErr, ok := exception.ErrorMapping[err.Error()]
		if !ok {
			myErr = exception.ErrorInternalServer
		}

		return api.NewResponse(
			api.WithMessage(err.Error()),
			api.WithError(myErr),
		).Send(ctx)
	}

	return api.NewResponse(
		api.WithHttpCode(http.StatusCreated),
		api.WithMessage("Register successfully"),
		api.WithData(map[string]interface{}{
			"id": result,
		}),
	).Send(ctx)
}

func (h authHandler) Login(ctx *fiber.Ctx) error {
    var req = loginUserSchema{}

	if err := ctx.BodyParser(&req); err != nil {
		myErr := exception.ErrorBadRequest
		return api.NewResponse(
			api.WithMessage(err.Error()),
			api.WithError(myErr),
			api.WithHttpCode(http.StatusBadRequest),
			api.WithMessage("Login fail"),
		).Send(ctx)
	}

	model, err := h.user.FindUserByEmail(ctx.UserContext(), req.Email)
	if  err != nil { 
		return api.NewResponse(
			api.WithMessage("Email not registered"),
			api.WithError(exception.ErrorNotFound),
		).Send(ctx)
	}

	token, err := h.auth.ValidateUserCredentials(ctx.UserContext(), model, req.Password)
    if err != nil {
		myErr, ok := exception.ErrorMapping[err.Error()]
		if !ok {
			myErr = exception.ErrorInternalServer
		}

		return api.NewResponse(
			api.WithMessage(err.Error()),
			api.WithError(myErr),
		).Send(ctx)
	}

	// cfg, _ := configs.LoadConfig()

	cookie := new(fiber.Cookie)
  	cookie.Name = "refresh_token"
  	cookie.Value = token.RefreshToken
  	cookie.Expires = time.Now().Add(7 *24 * time.Hour)
	cookie.Secure = true
	cookie.HTTPOnly = true


	ctx.Cookie(cookie)

	return api.NewResponse(
		api.WithHttpCode(http.StatusOK),
		api.WithMessage("Login successfully"),
		api.WithData(map[string]interface{}{
			"access_token": token.AccessToken,
			"id": model.Id,
			"email": model.Email,
			"username": model.Username,
		}),
	).Send(ctx)
}

func (h authHandler) Refresh(ctx *fiber.Ctx) error {
	id := ctx.Locals("id")

	req := tokenSchema{
		Id: id.(string),
	}

	token, err := h.auth.Refresh(ctx.UserContext(), req)
    if err != nil {
		myErr, ok := exception.ErrorMapping[err.Error()]
		if !ok {
			myErr = exception.ErrorInternalServer
		}

		return api.NewResponse(
			api.WithMessage(err.Error()),
			api.WithError(myErr),
		).Send(ctx)
	}

	return api.NewResponse(
		api.WithHttpCode(http.StatusOK),
		api.WithMessage("Login successfully"),
		api.WithData(map[string]interface{}{
			"access_token": token,
		}),
	).Send(ctx)
}

func (h authHandler) Logout(ctx *fiber.Ctx) error {
	cfg, _ := environtment.LoadConfig()
	ctx.Cookie(&fiber.Cookie{
        Name:     "refresh_token",
        Expires:  time.Now().Add(-(time.Hour * 2)),
        HTTPOnly: true,
        Domain: cfg.App.DomainName,
		SameSite: "lax",
    })

	return api.NewResponse(
		api.WithHttpCode(http.StatusOK),
		api.WithMessage("Logout successfully"),
	).Send(ctx)
}