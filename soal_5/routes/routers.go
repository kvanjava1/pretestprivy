package routes

import (
	"net/http"
	c "pretestprivy/soal_5/controllers"

	"github.com/gorilla/mux"
)

func HandlerRoute(router *mux.Router) {
	router.HandleFunc("/", c.Free).Methods(http.MethodGet)
	router.HandleFunc("/ping", c.Ping).Methods(http.MethodGet)
	router.HandleFunc("/time", c.Time).Methods(http.MethodGet)
}
