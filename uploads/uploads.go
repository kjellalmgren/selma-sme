package uploads

import "net/http"

// Upload documentation
func Upload(w http.ResponseWriter, r *http.Request) {
	//processid := r.Header.Get("X-process-Id")
	//fmt.Printf("swagger executed, processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	err := r.MultipartReader()
	if err =! nil {
		fmt.Fprintf(w, "Error file upload - %s", err)
			w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
}
