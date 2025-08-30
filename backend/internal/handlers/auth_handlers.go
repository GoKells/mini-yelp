package handlers

import (

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
	Error interface{} `json:"error"`
}

type SignUpInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	// CallbackURL string `json:"callbackURL"`
	// RememberMe bool `json:"rememberMe"`
}

func SignUp(c *fiber.Ctx) error {
	data := new(SignUpInput)

	// Parse JSON body into User struct
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Message: "Error parsing data",
			Data: nil,
			Error: fiber.Map{
				"code": "INTERNAL_ERROR",
				"detail": "there was an error parsing the signup data",
			},
		})
	}

	// Validate input
	if data.Name == "" || data.Email == "" || data.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Message: "missing sign up credentials",
			Data: nil,
			Error: fiber.Map{
				"code": "MISSING_CRENDTIALS",
				"detail": "email, name, and password are required",
			},
		})
	}

	// Add user data to the database and set verified to false
	// Create a magic token to send to user email

	resp := Response{
		Success: true,
		Message: "Welcom to mini yelp",
		Data: fiber.Map{
			"emali": data.Email,
			"name": data.Name,
			"email-sent": true,
		},
		Error: nil,
	}

	return c.Status(200).JSON(resp)
}

func Login(c *fiber.Ctx) error {
	data := new(LoginInput)

	// Parse JSON body into User struct
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Message: "Error parsing data",
			Data: nil,
			Error: fiber.Map{
				"code": "INTERNAL_ERROR",
				"detail": "there was an error parsing login data",
			},
		})
	}

	// Validate input
	if data.Email == "" || data.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Message: "missing login credentials",
			Data: nil,
			Error: fiber.Map{
				"code": "MISSING_CRENDTIALS",
				"detail": "email, and password are required",
			},
		})
	}

	resp := Response{
		Success: true,
		Message: "Welcome back",
		Data: fiber.Map{
			"emali": data.Email,
			"name": "John Doe",
			"session": uuid.New(),
		},
		Error: nil,
	}

	return c.Status(200).JSON(resp)
}

func Logout(c *fiber.Ctx) error {
	id:= c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Message: "Logout unsuccessful",
			Data: nil,
			Error: fiber.Map{
				"code": "MISSING_SESSION",
				"detail": "no session has been found",
			},
		})
	}

	resp := Response{
		Success: true,
		Message: "Logout successful",
		Data: fiber.Map{
			"revoked": true,
		},
		Error: nil,
	}

	return c.Status(200).JSON(resp)
}

func CreateMagicLink(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func ConsumeMagicLink(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func CreatePasswordResetLink(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func ConsumePasswordResetLink(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}