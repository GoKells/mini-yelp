package handlers

import (
	"backend/internal/service"

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

type PasswordUpate struct {
	VerifyPassword    string `json:"verify"`
	Password string `json:"password"`
	Token string `json:"token"`
	// CallbackURL string `json:"callbackURL"`
	// RememberMe bool `json:"rememberMe"`
}

type ResetPassword struct {
	Email string `json:"email"`
}

type VerifyToken struct {
	Token string `json:"token"`
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
	err := service.UserResgistrationEmail(data.Name, data.Email, "<magic-token-string>")
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(Response{
			Success: false,
			Message: "error sending registration email",
			Data: nil,
			Error: fiber.Map{
				"code": "INTERNAL_ERRORE",
				"detail": "there was an issue sending a registration confirmation email",
			},
		})
	}

	resp := Response{
		Success: true,
		Message: "check your email for verification link",
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

func UpdatePassword(c *fiber.Ctx) error {
	data := new(PasswordUpate)

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
	if data.Password == "" || data.VerifyPassword == "" || data.Token == "" {
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

	if data.Password != data.VerifyPassword {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Message: "password and password verification are not the same",
			Data: nil,
			Error: fiber.Map{
				"code": "INVALID_CRENDTIALS",
				"detail": "password and verify have to be the same",
			},
		})
	}

	resp := Response{
		Success: true,
		Message: "Welcome back",
		Data: fiber.Map{
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

func CreatePasswordResetLink(c *fiber.Ctx) error {
	data := new(ResetPassword)

	// Parse JSON body into User struct
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Message: "Error parsing data",
			Data: nil,
			Error: fiber.Map{
				"code": "INTERNAL_ERROR",
				"detail": "there was an error parsing password reset data",
			},
		})
	}

	// Validate input
	if data.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Message: "missing credentials",
			Data: nil,
			Error: fiber.Map{
				"code": "MISSING_CRENDTIALS",
				"detail": "email is required",
			},
		})
	}

	// Check if account with email exists

	resp := Response{
		Success: true,
		Message: "string",
		Data: fiber.Map{
			"to": data.Email,
		},
		Error: nil,
	}

	return c.JSON(resp)
}

func VerifyPasswordResetLink(c *fiber.Ctx) error {
	data := new(VerifyToken)

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
	if data.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Message: "missing credentials",
			Data: nil,
			Error: fiber.Map{
				"code": "MISSING_CRENDTIALS",
				"detail": "email is required",
			},
		})
	}

	resp := Response{
		Success: true,
		Message: "string",
		Data: fiber.Map{
			"sent": true,
		},
		Error: nil,
	}

	return c.JSON(resp)
}

func VerifySignupLink(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}
