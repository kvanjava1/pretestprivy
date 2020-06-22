package main

import (
	c "soal_5/controllers"

	"github.com/gorilla/mux"
)

func handlerRoute(r *mux.Router) {
	r.HandleFunc("/", c.Home)
}
