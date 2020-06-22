package main

import (
	"log"
	"net/http"
	router "pretestprivy/soal_5/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	router.HandlerRoute(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
