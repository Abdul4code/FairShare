package main

import (
	"net/http"

	"github.com/Abdul4code/FairShare/internal"
	"github.com/julienschmidt/httprouter"
)

// Router constructs and returns the application's HTTP router with routes and custom
// NotFound and MethodNotAllowed handlers wired up.
func (app *application) Router() *httprouter.Router {
	// instantiate new router
	router := httprouter.New()

	// customize the  not found and method not allowed page
	router.NotFound = http.HandlerFunc(internal.NotFoundError)
	router.MethodNotAllowed = http.HandlerFunc(internal.MethodNotAllowed)

	// health check route
	router.HandlerFunc(http.MethodGet, "/v1/health", app.healthCheckHandler)

	// groups routes
	router.HandlerFunc(http.MethodPost, "/v1/groups", app.CreateGroupHandler)
	router.HandlerFunc(http.MethodGet, "/v1/groups/:id", app.GetGroupHandler)
	router.HandlerFunc(http.MethodPut, "/v1/groups/:id", app.UpdateGroupHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/groups/:id", app.DeleteGroupHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/groups/:id", app.PatchGroupHandler)
	router.HandlerFunc(http.MethodGet, "/v1/groups", app.GetGroupsHandler)

	return router
}
