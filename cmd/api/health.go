package main

import (
	"net/http"

	"github.com/Abdul4code/FairShare/internal"
)

// create a health check handler
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	message := map[string]string{"status": "available", "environment": app.Config.Env}
	internal.WriteJSON(w, http.StatusOK, message)
}
