package users

import "github.com/gofiber/fiber/v2"

// Handler interface is
type Handler interface {
	SignUp() fiber.Handler
}
