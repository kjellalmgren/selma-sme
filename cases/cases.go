package cases

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Cases struct
type Cases struct {
	CaseID string `json:"caseId,omitempty"`
}

//
//	OperationId: ReserverCaseId
func ReserveCaseId(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("ReserveCaseID executed...\r\n")
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//
	vars := mux.Vars(r)
	fmt.Println(vars["processId"])    // processId
	fmt.Println(vars["customerId"])   // customerId
	fmt.Println(vars["caseIdStatus"]) // caseIdStatus
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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//
	vars := mux.Vars(r)
	fmt.Println(vars["processId"])    // processId
	fmt.Println(vars["caseId"])       // caseId
	fmt.Println(vars["caseIdStatus"]) // caseIdStatus
	//
	w.WriteHeader(http.StatusOK)
}
