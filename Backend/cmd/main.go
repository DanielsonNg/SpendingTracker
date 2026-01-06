package main

import (
	"log"

	"github.com/DanielsonNg/SpendingTracker/tree/main/Backend/cmd/api"
	"github.com/DanielsonNg/SpendingTracker/tree/main/Backend/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Environment variable not found")
	}
	db, err := database.NewPostgres()
	if err != nil {
		log.Fatal("Database Connection Failed", err)
	}
	defer db.Close()

	log.Println("Connected to database")
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
