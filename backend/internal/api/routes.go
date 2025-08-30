package api

import (
	"backend/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (a *FiberApi) RegisterFiberRoutes() {
	// Apply CORS middleware
	a.App.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Accept,Authorization,Content-Type",
		AllowCredentials: false, // credentials require explicit origins
		MaxAge:           300,
	}))

	auth := a.Group("/auth")
	a.RegisterAuthRoutes(auth)

}

func (a *FiberApi) RegisterAuthRoutes(router fiber.Router) {
	router.Post("/signup", handlers.SignUp)
	router.Post("/login", handlers.Login)
	router.Delete("/logout/:id", handlers.Logout)
	router.Delete("/verify/:token", handlers.ConsumeMagicLink)
}
