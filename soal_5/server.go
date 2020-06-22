package main

import (
	"log"
	"net/http"
	hr "pretestprivy/soal_5/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	hr.HandlerRoute(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
