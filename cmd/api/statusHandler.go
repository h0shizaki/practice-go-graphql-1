package main

import (
	"net/http"
)

func (app *Application) statusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := Status{
		Status:      "Available",
		Environment: app.Config.Environment,
		Version:     app.Config.Version,
	}

	app.writeJSON(w, 200, currentStatus, "status")
	// js, err := json.MarshalIndent(currentStatus, "", "\t")

	// if err != nil {
	// 	app.Logger.Println(err)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(js)

}
