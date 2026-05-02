package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rjmcnamara10/forage/db"
	"github.com/rjmcnamara10/forage/db/repository"
	"github.com/rjmcnamara10/forage/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := db.NewConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close()

	repos := repository.NewRepositories(conn)

	router := gin.Default()
	handlers.RegisterRoutes(router, repos)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
