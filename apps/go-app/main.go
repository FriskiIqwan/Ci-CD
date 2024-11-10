package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"go-app/internal/routes"
)

// Config holds the application configuration
type Config struct {
	ListenAddress string
}

func loadConfig() Config {
	listenAddress := os.Getenv("LISTEN_ADDRESS")
	if listenAddress == "" {
		listenAddress = "127.0.0.1:8080" // default address
	}
	return Config{
		ListenAddress: listenAddress,
	}
}

func main() {
	// Load configuration
	config := loadConfig()

	// Initialize Echo instance
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.INFO)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:8000",
			"https://example.com",
		},
		AllowMethods: []string{
			http.MethodGet, http.MethodHead, http.MethodPut,
			http.MethodPatch, http.MethodPost, http.MethodDelete,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			"X-Client-Info",
		},
	}))

	// Register routes
	routes.SetupRoutes(e)

	// Create a channel to listen for interrupt or terminate signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Start server in a goroutine
	go func() {
		e.Logger.Infof("Starting server on %s", config.ListenAddress)
		if err := e.Start(config.ListenAddress); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatalf("Shutting down the server: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-quit
	e.Logger.Info("Shutting down server...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatalf("Server forced to shutdown: %v", err)
	}

	e.Logger.Info("Server exiting")
}
