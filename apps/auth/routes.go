package auth

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/shelllbyyyyy/belajar-api-go/internal/api"
)

func UserRoute(router fiber.Router, db *sql.DB) {
	repo := newUserRepository(db)
	usecase := NewUserUseCase(repo)
	handler := newUserHandler(*usecase)

	_ = handler

	authRouter := router.Group("/api/v1/users")
	{
		authRouter.Get("/:email", api.CheckAuth(), handler.findByEmail)
		authRouter.Get("/:id", api.CheckAuth(), handler.findById)
		authRouter.Patch("/", api.CheckAuth(), handler.update)
	}
}

func AuthRoute(router fiber.Router, db *sql.DB) {
	repo := newUserRepository(db)
	userUsecase := NewAuthUseCase(repo)
	authUsecase := NewUserUseCase(repo)
	handler := newAuthHandler(*userUsecase, *authUsecase)

	_ = handler

	authRouter := router.Group("/api/v1/auth")
	{
		authRouter.Post("/login",  handler.Login)
		authRouter.Post("/register",  handler.Register)
		authRouter.Post("/refresh", api.RefreshToken(), handler.Refresh)
		authRouter.Post("/logout", api.CheckAuth(), handler.Logout)
	}
}