package auth

import (
	"net/http"

	"github.com/koesterjannik/starter/logger"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Send the text ok back to the client
	logger.Logger.Info("Health check performed")
	w.Write([]byte("ok"))

}
