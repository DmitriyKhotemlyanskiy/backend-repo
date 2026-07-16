package main

import (
	"log"
	"net/http"

	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Load configuration
	cfg := config.LoadConfig()

	// 2. Connect to database
	log.Printf("Connecting to MongoDB at %s...", cfg.MongoURI)
	db, err := database.ConnectDB(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	log.Println("Successfully connected to MongoDB")

	// 3. Setup Gin router and routes
	r := gin.Default()

	// CORS Middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Liveness and Readiness probes for Kubernetes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	// Initialize handlers
	h := handler.NewHandler(db)

	// API Endpoints
	r.GET("/api/hotels", h.GetHotels)
	r.POST("/api/reservations", h.CreateReservation)
	r.GET("/api/reservations/lookup", h.LookupReservation)
	r.DELETE("/api/reservations/:id", h.CancelReservation)

	// Start server
	log.Printf("Server is running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
