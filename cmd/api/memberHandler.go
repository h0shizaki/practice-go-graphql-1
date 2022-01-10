package main

import "net/http"

func (app *Application) getAllMember(w http.ResponseWriter, r *http.Request) {
	member, err := app.Models.DB.All()

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, member, "all-member")

	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
