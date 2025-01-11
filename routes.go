package main

import(
	"net/http"
)

// This struct is for data structures of routes
type Route struct {
	Name string // Name of the route. E.g. "sign-in"
	Method string // Method of the route. E.g. GET, POST
	HandlerFunc http.HandlerFunc // a handler function
	Pattern string // URL path pattern. E.g. "/products" or "/products/{prdId}
}

// slice a "Route" struct that allows grouping multiple "Route" structs together
type Routes []Route

// 
var routes = Routes{
	Route{
		"Index",
		"GET",
		Index,
		"/",
	},
	Route{
		"TodolistIndex",
		"GET",
		TodolistIndex,
		"/TodolistIndex",
	},
	Route{
		"TodolistShow",
		"GET",
		TodolistShow,
		"/TodolistShow",
	},
	Route{
		"TodolistCreate",
		"POST",
		TodolistCreate,
		"/TodolistCreate",
	},
}