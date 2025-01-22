package auth

import (
	"net/http"

	"github.com/shelllbyyyyy/belajar-api-go/internal/api"
	"github.com/shelllbyyyyy/belajar-api-go/internal/exception"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userUsecase UserUseCase
}

func newUserHandler(uc UserUseCase) userHandler {
	return userHandler{
		userUsecase: uc,
	}
}

func (h userHandler) findByEmail(ctx *fiber.Ctx) error {
	var param = ctx.Params("email")

	model, err := h.userUsecase.FindUserByEmail(ctx.UserContext(), param)
	if err != nil {
		return api.NewResponse(
			api.WithMessage("User not found"),
			api.WithError(exception.ErrorNotFound),
		).Send(ctx)
	}

	response := toUserResponse(model)

	return api.NewResponse(
		api.WithHttpCode(http.StatusOK),
		api.WithMessage("User Found"),
		api.WithData(response),
	).Send(ctx)
}

func (h userHandler) findById(ctx *fiber.Ctx) error {
	var param = ctx.Params("id")

	model, err := h.userUsecase.FindUserById(ctx.UserContext(), param)
	if err != nil {
		return api.NewResponse(
			api.WithMessage("User not found"),
			api.WithError(exception.ErrorNotFound),
		).Send(ctx)
	}

	response := toUserResponse(model)

	return api.NewResponse(
		api.WithHttpCode(http.StatusOK),
		api.WithMessage("User Found"),
		api.WithData(response),
	).Send(ctx)
}

func (h userHandler) update(ctx *fiber.Ctx) error {
	var param = ctx.Locals("id").(string)
	var req = updateUserSchema{}

	if err := ctx.BodyParser(&req); err != nil {
		myErr := exception.ErrorBadRequest
		return api.NewResponse(
			api.WithMessage(err.Error()),
			api.WithError(myErr),
			api.WithHttpCode(http.StatusBadRequest),
			api.WithMessage("register fail"),
		).Send(ctx)
	}

	model, err := h.userUsecase.FindUserById(ctx.UserContext(), param)
	if err != nil {
		return api.NewResponse(
			api.WithMessage("User not found"),
			api.WithError(exception.ErrorNotFound),
		).Send(ctx)
	}

	result, err := h.userUsecase.Update(ctx.UserContext(), model, req)
	if err != nil {
		return api.NewResponse(
			api.WithMessage("Update user failed"),
			api.WithError(exception.ErrorBadRequest),
		).Send(ctx)
	}

	return api.NewResponse(
		api.WithHttpCode(http.StatusOK),
		api.WithMessage("Update user success"),
		api.WithData(result),
	).Send(ctx)
}