package server

import (
	"github.com/gofiber/fiber/v2"
	userHandler "github.com/jumayevgadam/todo_app-fiber/internal/users/delivery/handler"
	userRoutes "github.com/jumayevgadam/todo_app-fiber/internal/users/routes"
	userService "github.com/jumayevgadam/todo_app-fiber/internal/users/usecase"
)

// EndpointsURL are
const (
	userGroup = "/api/v1/user"
)

func (s *Server) MapHandlers(f *fiber.App) {
	// Init Services
	UserService := userService.NewUserService(s.dataStore)

	// InitHandlers
	UserHandler := userHandler.NewUserHandler(UserService)

	// Init MainGroup Routes
	UserGroup := s.fiber.Group(userGroup)
	// Init Mapping Routes
	userRoutes.MapUserRoutes(UserGroup, UserHandler)
}
