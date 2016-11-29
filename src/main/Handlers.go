package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

/**
 * JSON PING response
 */
func PingHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("pong")
}

func IdHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    json.NewEncoder(w).Encode(id)
}

func AddBeer(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")
		salcohol := r.FormValue("alcohol")//return string
		alcohol, err := strconv.ParseFloat(salcohol, 32)
		if err != nil {
		  // handle the error in some way
			http.Error(w, http.StatusText(400),
                400)
		} else {
			log.Print("adding beer " + name)
			b := Beer{Name: "Sean", Alcohol: float32(alcohol)}
			addBeer(b)
		}
}

func GetBeers(w http.ResponseWriter, r *http.Request) {
		log.Print("get beers request")
    fmt.Fprintf(w, getBeers())
}
