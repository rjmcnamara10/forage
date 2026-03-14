package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rjmcnamara10/forage/db"
	"github.com/rjmcnamara10/forage/db/repository"
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

	// Initialize repositories
	repos := repository.NewRepositories(conn)

	// TODO: Initialize Gin router and add handlers
	// For now, just verify the repositories are initialized
	log.Println("Repositories initialized successfully")
	log.Printf("Units repository: %v\n", repos.Units)
	log.Printf("Items repository: %v\n", repos.Items)
	log.Printf("Meals repository: %v\n", repos.Meals)
	log.Printf("Stores repository: %v\n", repos.Stores)
	log.Printf("Shopping lists repository: %v\n", repos.ShoppingLists)
}
