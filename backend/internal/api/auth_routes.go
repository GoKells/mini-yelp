package api

import (
	"backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *FiberApi) RegisterAuthRoutes(router fiber.Router) {
	router.Post("/signup", handlers.SignUp) // Sends verification email to activate account
	router.Post("/login", handlers.Login) // Login active user
	router.Delete("/logout", handlers.Logout) // logout user
	router.Post("/password/reset", handlers.CreatePasswordResetLink) // send link to reset password
	router.Post("/password/new", handlers.UpdatePassword) // receives the new password
	router.Post("/verify/reset", handlers.VerifyPasswordResetLink) // verifies token to update password and returns a able to reset token
	router.Post("/verify/signup", handlers.VerifySignupLink) // verifies token to activate user and automatically logs in user
}
