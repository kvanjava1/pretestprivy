package routes

import (
	c "pretestprivy/soal_5/controllers"

	"github.com/gorilla/mux"
)

func HandlerRoute(r *mux.Router) {
	r.HandleFunc("/", c.Home)
}
