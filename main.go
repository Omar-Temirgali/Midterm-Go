package main

import (
	"log"
	"net/http"

	"github.com/Omar-Temirgali/go-service/routes"
)

func main() {
	router := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":5000", router))
}
