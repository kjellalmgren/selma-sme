package cases

import (
	"encoding/json"
	"fmt"
	"net/http"
	"selmasme/models"
)

//
//	OperationId: ReserverCaseId
func ReserveCaseID(w http.ResponseWriter, r *http.Request) {

	//vars := mux.Vars(r)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID, caseIdStatus")
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	processid := r.Header.Get("X-process-Id")
	//
	//for k, v := range r.Header {
	//	fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
	//	fmt.Println("Header field %q, Value %q\n", k, v)
	//}
	fmt.Printf("ReserveCaseID executed...X-Process-ID: %s\r\n", processid)
	//
	var cases models.Cases
	cases.CaseID = "108000" // Hardcoded
	//
	if err := json.NewEncoder(w).Encode(cases); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

//
// SetCaseIdStatus
func SetCaseIDStatus(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("SetCaseIDStatus executed...\r\n")
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "PATCH, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID, caseIdStatus")
	//
	processid := r.Header.Get("X-process-ID")
	fmt.Printf("setCaseIDStatus executed...X-Process-ID: %s\r\n", processid)
	//
	//
	w.WriteHeader(http.StatusOK)
}
