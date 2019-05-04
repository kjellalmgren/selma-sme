package swagger

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// swagger documentation, this holds the end-point to API definitions
func Swagger(w http.ResponseWriter, r *http.Request) {

	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("swagger executed, processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	file, err := ioutil.ReadFile("json/selma-en-v0.6.0.yaml")
	if err != nil {
		fmt.Fprintf(w, "Error reading selma-en-v0.6.0.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(file)
	w.WriteHeader(http.StatusOK)
	//fmt.Fprint(os.Stderr, fmt.Sprintf("selma-en-v0.5.9"))
}
