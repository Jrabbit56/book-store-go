package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jrabbit56/book-store/internal/core/domain"
	"github.com/jrabbit56/book-store/internal/core/ports"
)

type UserHandler struct {
	service ports.UserService
}

func NewUserHandler(service ports.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// RegisterUser godoc
//
//	@Summary		Register a new user
//	@Description	Registers a new user in the system.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body	domain.SwaggerUser	true	"User Registration"
//	@Success		200		{object}	map[string]string
//	@Router			/register [post]
func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	user := new(domain.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.RegisterUser(user); err != nil {
		return c.JSON(fiber.Map{"message": "Email already exit!!"})
	}
	return c.JSON(fiber.Map{"message": "Register Successfully!!"})
}

// LoginUser godoc
//
//	@Summary		Authenticate a user
//	@Description	Logs in a user, returns a JWT token, and sets it as an HTTP-only cookie
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		domain.User			true	"User Credentials"
//	@Success		200		{object}	map[string]string	"message":"Login Successfully!!"
//	@Failure		400		{object}	map[string]string	"error":"Bad Request Error Message"
//	@Failure		401		{object}	map[string]string	"error":"Unauthorized Error Message"
//	@Header			200		{string}	Set-Cookie			"Contains the JWT token"
//	@Router			/login [post]
func (h *UserHandler) LoginUser(c *fiber.Ctx) error {
	user := new(domain.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := h.service.LoginUser(user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	//Keep Cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24), // 24 hours
		HTTPOnly: true,                           // Only accessible via HTTP(S) requests
	})

	return c.JSON(fiber.Map{"message": "Login Successfully!!"})
	
}
