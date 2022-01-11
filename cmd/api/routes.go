package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Create route
func (app *Application) routes() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/member", app.getAllMember)
	router.HandlerFunc(http.MethodGet, "/v1/member/:id", app.getOneMember)

	return router
}
