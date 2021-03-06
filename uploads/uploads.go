package uploads

import (
	"fmt"
	"net/http"
)

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
	//MultipartReader r1
	//body := &bytes.Buffer{}
	//rbody := r.Body.Read
	//writer := multipart.NewWriter(body)
	//r.Body.Close()
	//
	r2, err := r.MultipartReader()
	if r2 != nil {
		fmt.Fprintf(w, "Error when using file upload - %d", err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
}
