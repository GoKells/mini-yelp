package api

import (

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
