package helpers

import (
	"fmt"
	"net/http"
	"strconv"
)

func Success(w http.ResponseWriter, resPayload string, debugMode bool) {
	resStatus := http.StatusOK

	if debugMode {
		fmt.Print("status code : " + "\n" + strconv.Itoa(resStatus) + "\n")
		fmt.Print("payload : " + "\n" + resPayload + "\n")
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(resStatus)
	w.Write([]byte(resPayload))
}
