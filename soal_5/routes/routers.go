package routes

import (
	"net/http"
	cExample "pretestprivy/soal_5/controllers/example"

	"github.com/gorilla/mux"
)

func HandlerRoute(router *mux.Router) {
	router.HandleFunc("/", cExample.Free).Methods(http.MethodGet)
	router.HandleFunc("/ping", cExample.Ping).Methods(http.MethodGet)
	router.HandleFunc("/time", cExample.Time).Methods(http.MethodGet)
}
