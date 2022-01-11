package main

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

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

func (app *Application) getOneMember(w http.ResponseWriter, r *http.Request) {
	param := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(param.ByName("id"))

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	member, err := app.Models.DB.One(id)

	err = app.writeJSON(w, http.StatusOK, member, "member")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
