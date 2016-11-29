package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
  "flag"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Ping", "GET", "/ping", PingHandler},
  Route{"IdHandler", "GET", "/id/{id}", IdHandler},
	Route{"AddBeer", "POST", "/beers", AddBeer},
	Route{"GetBeer", "GET", "/beers", GetBeers},
}

func NewRouter() *mux.Router {
    var dir string
    flag.StringVar(&dir, "dir", "./src/pages/", "./src/pages/")
    flag.Parse()
    log.Printf(dir)

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		log.Printf("routing : " + route.Pattern)
		router.
			Methods(route.Method).Path(route.Pattern).
			Name(route.Name).Handler(route.HandlerFunc)
	}
    // static page
    var d = http.StripPrefix("/pages/", http.FileServer(http.Dir(dir)))
    //log.Printf("routing pages: " + d)
    router.PathPrefix("/pages/").Handler(d)

	return router
}
