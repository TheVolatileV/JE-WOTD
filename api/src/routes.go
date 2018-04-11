package main

import "net/http"

// Route defines our structure for routes within our router
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is an array of Route
type Routes []Route

var routes = Routes{
	Route{
		"GetWord",
		"GET",
		"/api/v1",
		handler,
	},
	Route{
		"GetWordForce",
		"GET",
		"/api/v1/force",
		forceNewWordHandler,
	},
	Route{
		"Register",
		"POST",
		"/api/v1/register",
		insertEmailHandler,
	},
}
