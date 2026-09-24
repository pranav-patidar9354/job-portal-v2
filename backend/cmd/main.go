package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/pranav-patidar9354/job-portal-v2/backend/config"
	"github.com/pranav-patidar9354/job-portal-v2/backend/internal/routes"
)

func main() {
	_ = godotenv.Load()

	if err := config.ConnectDatabase(); err != nil {
		log.Fatal("database connection failed: ", err)
	}

	router := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Job Portal V2 API running on http://localhost:" + port)
	log.Fatal(router.Run(":" + port))
}
