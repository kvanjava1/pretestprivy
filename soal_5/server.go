package main

import (
	"fmt"
	"log"
	"net/http"
	router "pretestprivy/soal_5/routes"
	"strconv"

	dotenv "pretestprivy/soal_5/helpers/dotenv"

	"github.com/gorilla/mux"
)

var config = dotenv.GetConfigValues()

func main() {
	r := mux.NewRouter()
	router.HandlerRoute(r)
	debugMode, _ := strconv.ParseBool(config["DEBUG_MODE"])

	if debugMode {
		fmt.Print("server start " + config["APP_NAME"] + " in port " + config["PORT"] + " (" + config["ENV"] + ")\n")
	}

	log.Fatal(http.ListenAndServe(":"+config["PORT"], r))

}
