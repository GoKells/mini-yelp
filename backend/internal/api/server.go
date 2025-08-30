package api

import (
	"github.com/gofiber/fiber/v2"
)

type FiberApi struct {
	*fiber.App
}



func New() *FiberApi {
	api := &FiberApi{
		App: fiber.New(fiber.Config{
			ServerHeader: "mini-yelp-api",
			AppName:      "mini-yelp-api",
		}),
	}

	return api
}
