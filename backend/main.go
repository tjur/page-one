package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Port int `env:"BACKEND_PORT" envDefault:"8080"`
}

func main() {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		log.Fatalf("%+v", err)
	}

	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Define a simple GET endpoint
	router.GET("/ping", func(c *gin.Context) {
		t := time.Now()
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"time":    t.Format("2006-01-02 15:04:05"),
		})
	})

	// Start server
	router.Run(fmt.Sprintf(":%d", config.Port))
}
