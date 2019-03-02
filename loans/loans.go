package loans

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Loans
type Loan struct {
	ProcessID     string `json:"processId,omitempty"`
	LoanID        string `json:"loanId,omitempty"`
	LoanNumber    string `json:"loanNumber,omitempty"`
	LoanAmount    string `json:"loanAmount,omitempty"`
	NumberOfLoans string `json:"numberOfLoans,omitempty"`
	PurposeOfLoan string `json:"purposeOfLoan,omitempty"`
}

// GetLoans
func GetLoans(w http.ResponseWriter, r *http.Request) {

	var loans []Loan
	//
	vars := mux.Vars(r)
	fmt.Printf("getLoans executed: processId: %s...\r\n", vars["processId"])
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//
	switch vars["processId"] {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		loans = append(loans,
			Loan{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				LoanID:        "9b8e4822-3cb7-11e9-b210-d663bd873d93",
				LoanNumber:    "930101011212",
				LoanAmount:    "2000000",
				NumberOfLoans: "1",
				PurposeOfLoan: "FAF"})
		loans = append(loans,
			Loan{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				LoanID:        "76f0b120-3cb8-11e9-b210-d663bd873d93",
				LoanNumber:    "930101011213",
				LoanAmount:    "200000",
				NumberOfLoans: "1",
				PurposeOfLoan: "RTM"})
	}
	//
	if err := json.NewEncoder(w).Encode(loans); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}
