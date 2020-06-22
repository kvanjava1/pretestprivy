package responses

import (
	"fmt"
	"net/http"
	dotenv "pretestprivy/soal_5/helpers/dotenv"
	"strconv"
)

var config = dotenv.GetConfigValues()

func Success(w http.ResponseWriter, resPayload string) {
	resStatus := http.StatusOK
	debugMode, _ := strconv.ParseBool(config["DEBUG_MODE"])

	if debugMode {
		fmt.Print("status code : " + "\n" + strconv.Itoa(resStatus) + "\n")
		fmt.Print("payload : " + "\n" + resPayload + "\n")
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(resStatus)
	w.Write([]byte(resPayload))
}
