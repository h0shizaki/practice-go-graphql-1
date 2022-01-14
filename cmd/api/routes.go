package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Create route
func (app *Application) routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/member", app.getAllMember)
	router.HandlerFunc(http.MethodGet, "/v1/member/:id", app.getOneMember)

	router.HandlerFunc(http.MethodPost, "/v1/graphql", app.memberGraphQL)

	return app.enableCOR(router)
}
