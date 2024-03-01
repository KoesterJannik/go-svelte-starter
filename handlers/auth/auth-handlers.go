package auth

import (
	"encoding/json"
	"net/http"

	"github.com/koesterjannik/starter/db"
	"github.com/koesterjannik/starter/logger"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Send the text ok back to the client
	logger.Logger.Info("Health check performed")
	w.Write([]byte("ok"))

}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if users == nil {
		// return json with empty array
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))
		return
	}

	logger.Logger.Info("Get all users performed")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
