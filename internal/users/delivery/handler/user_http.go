package handler

import (
	"github.com/gofiber/fiber/v2"
	userModel "github.com/jumayevgadam/todo_app-fiber/internal/models/users"
	userServInterface "github.com/jumayevgadam/todo_app-fiber/internal/users"
	"github.com/jumayevgadam/todo_app-fiber/pkg/errlist"
	"github.com/jumayevgadam/todo_app-fiber/pkg/utils"
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
		if err := utils.ReadRequest(c, &requestForSignUp); err != nil {
			return c.JSON(errlist.Response(err))
		}

		data, err := f.service.SignUp(c.Context(), &requestForSignUp)
		if err != nil {
			return c.JSON(errlist.Response(err))
		}

		return c.Status(200).JSON(data)
	}
}
