package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/koesterjannik/starter/db"
	"github.com/koesterjannik/starter/handlers/auth"
	"github.com/koesterjannik/starter/logger"
)

func main() {
	godotenv.Load()
	logger.InitLogger()
	db.ConnectToDb()
	server := http.NewServeMux()
	// create a health check route
	server.HandleFunc("GET /health", auth.HealthCheck)
	server.HandleFunc("GET /users", auth.GetAllUsersHandler)

	val := os.Getenv("PORT")
	if val == "" {
		panic("Environment variable ENV not set.")
	}
	// start the server
	http.ListenAndServe(":"+val, server)

}
