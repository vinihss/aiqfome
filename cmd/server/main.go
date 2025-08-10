package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vinihss/aiqfome/config"
	_ "github.com/vinihss/aiqfome/docs"
	"github.com/vinihss/aiqfome/internal/infrastructure/database/models"
	"github.com/vinihss/aiqfome/internal/routes"
	"log"
	"os"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Run() error {
	config.ConnectDB()
	err := config.DB.AutoMigrate(&models.Favorite{}, &models.Customer{})
	if err != nil {
		return errors.New(fmt.Sprintf("Error migrating database: %v", err))
	}

	// Set up Gin router
	r := gin.Default()
	routes.SetupRoutes(r)

	err = r.Run(s.addr)
	if err != nil {
		return errors.New(fmt.Sprintf("Error starting server: %v", err))
	}
	// Start the server
	return nil
}

// @title Favorite API
// @version 1.0
// @description API para favoritar produtos
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	server := NewServer(":8080")
	err := server.Run()
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
		os.Exit(1)
	}
}
