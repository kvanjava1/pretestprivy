package example

import (
	"net/http"
	response "pretestprivy/soal_5/helpers/responses"
	"time"
)

func Free(w http.ResponseWriter, r *http.Request) {
	response.Success(w, "hasil jawaban nomer 5")
}

func Ping(w http.ResponseWriter, r *http.Request) {
	response.Success(w, "pong")
}

func Time(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	response.Success(w, currentTime.Format("2006-01-02 15:04:05"))
}
