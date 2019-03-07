package cases

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Cases struct
type Cases struct {
	CaseID string `json:"caseId"`
}

//
//	OperationId: ReserverCaseId
func ReserveCaseID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	varsh := r.Header
	//
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
		fmt.Println("Header field %q, Value %q\n", k, v)
	}
	fmt.Printf("kjell-header: %s\r\n", varsh["x-process-id"])
	fmt.Printf("ReserveCaseID executed...ProcessId: %s /customerID: %s / caseIdStatus: %s\r\n",
		vars["processId"], vars["customerId"], vars["caseIdStatus"])
	//
	var cases Cases
	cases.CaseID = "108000" // Hardcoded
	//
	if err := json.NewEncoder(w).Encode(cases); err != nil {
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
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	vars := mux.Vars(r)
	fmt.Println(vars["processId"])    // processId
	fmt.Println(vars["caseId"])       // caseId
	fmt.Println(vars["caseIdStatus"]) // caseIdStatus
	//
	w.WriteHeader(http.StatusOK)
}
