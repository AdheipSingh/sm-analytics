package main

import (
	"log"
	"net/http"

	"sm-analytics/routes"
)

func main() {

	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))

}
