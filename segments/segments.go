package segments

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Segments documentation, this holds the end-point to API definitions
func Segments(w http.ResponseWriter, r *http.Request) {

	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("Segments executed, processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	file, err := ioutil.ReadFile("json/segment_training.csv")
	if err != nil {
		fmt.Fprintf(w, "Error reading segment_training.csv - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(file)
	w.WriteHeader(http.StatusOK)
}
