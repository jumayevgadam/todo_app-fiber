package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/todo_app-fiber/internal/users"
)

// MapUserRoutes is
func MapUserRoutes(userGroup fiber.Router, userHandler users.Handler) {
	// Routes are
	userGroup.Post("/sign-up", userHandler.SignUp())
}
