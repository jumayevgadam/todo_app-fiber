package server

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/jumayevgadam/todo_app-fiber/docs" // Swagger generated docs
	"github.com/jumayevgadam/todo_app-fiber/internal/config"
	"github.com/jumayevgadam/todo_app-fiber/internal/database"
	fiberswagger "github.com/swaggo/fiber-swagger"
)

// Server struct is
type Server struct {
	fiber     *fiber.App
	cfg       *config.Config
	dataStore database.DataStore
}

// NewServer is
func NewServer(
	cfg *config.Config,
	dataStore database.DataStore,
) *Server {
	server := &Server{
		fiber:     fiber.New(),
		cfg:       cfg,
		dataStore: dataStore,
	}

	return server
}

// Run is
func (s *Server) Run() error {
	s.MapHandlers(s.fiber)
	s.fiber.Get("/swagger/*", fiberswagger.WrapHandler)
	return s.fiber.Listen(s.cfg.Server.HTTPPort)
}
