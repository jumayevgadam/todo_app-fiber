package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/todo_app-fiber/internal/config"
	"github.com/jumayevgadam/todo_app-fiber/internal/connection"
)

func main() {
	// LoadConfig is
	_, err := config.LoadConfig()
	if err != nil {
		log.Printf("main.LoadConfig: %v", err.Error())
	}

	mysqlDB, err := connection.NewDBConnection(context.Background(), config.MySQL{})
	if err != nil {
		log.Printf("errror in db connection: %v", err.Error())
	}
	log.Println("Successfully connected to mysqlDB")

	defer func() {
		if mysqlDB != nil {
			if err := mysqlDB.Close(); err != nil {
				log.Printf("[main][mysqlDB.Close]: %v", err.Error())
			}
		}
	}()
	app := fiber.New()
	app.Get("/post", func(c *fiber.Ctx) error {
		return c.SendString("Hello EveryOne")
	})

	app.Listen(":8080")

	// Handle SIGINT and SIGTERM for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down...")
}
