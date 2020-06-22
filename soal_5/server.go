package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	handlerRoute(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
