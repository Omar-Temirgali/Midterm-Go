package routes

import (
	"net/http"

	"github.com/Omar-Temirgali/go-service/config"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		config.Index,
	},
	Route{
		"Fetching all",
		"GET",
		"/all",
		config.ShowAll,
	},
	Route{
		"Show",
		"GET",
		"/store/{key}",
		config.Show,
	},
	Route{
		"Update or Insert",
		"PUT",
		"/store/{key}/{value}",
		config.UpdateAndInsert,
	},
}
