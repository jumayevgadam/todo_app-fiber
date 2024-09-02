package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/todo_app-fiber/internal/config"
	"github.com/jumayevgadam/todo_app-fiber/internal/database"
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
	return s.fiber.Listen(s.cfg.Server.HTTPPort)
}
