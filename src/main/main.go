package main

import (
	"net/http"
  "log"
)

func main() {
	log.Println("Starting app.\n")
    //start webserver
    router := NewRouter()
    http.ListenAndServe(":8080", router)
    log.Fatal()
}
