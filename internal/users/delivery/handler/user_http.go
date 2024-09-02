package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	userModel "github.com/jumayevgadam/todo_app-fiber/internal/models/users"
	userServInterface "github.com/jumayevgadam/todo_app-fiber/internal/users"
)

// UsersHandler is
type UsersHandler struct {
	service userServInterface.Service
}

// NewUserHandler is
func NewUserHandler(service userServInterface.Service) *UsersHandler {
	return &UsersHandler{service: service}
}

// SignUp is
func (f *UsersHandler) SignUp() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestForSignUp userModel.SignUpReq
		if err := c.BodyParser(&requestForSignUp); err != nil {
			return c.Context().Err()
		}

		data, err := f.service.SignUp(c.Context(), &requestForSignUp)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.Status(200).JSON(data)
	}
}
