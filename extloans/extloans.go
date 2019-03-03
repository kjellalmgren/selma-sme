package extloans

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// ExtLoans
type ExtLoan struct {
	ProcessID         string  `json:"processId,omitempty"`
	CustomerID        string  `json:"customerId,omitempty"`
	ExtLoanID         string  `json:"extloanId,omitempty"`
	ExtCreditInstitut string  `json:"extCreditInstitut,omitempty"`
	ExtLoanClearing   string  `json:"extLoanClearing,omitempty"`
	ExtLoanNumber     string  `json:"extloanNumber,omitempty"`
	ExtLoanAmount     float32 `json:"extLoanAmount,omitempty"`
	ExtRedeemLoan     bool    `json:"extRedeemLoan"`
}

// GetExtLoans
func GetExtLoans(w http.ResponseWriter, r *http.Request) {

	var extloans []ExtLoan
	//
	vars := mux.Vars(r)
	fmt.Printf("getExtLoans executed: processId: %s...\r\n", vars["processId"])
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//
	switch vars["processId"] {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		extloans = append(extloans,
			ExtLoan{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CustomerID:        "19640120-3887",
				ExtLoanID:         "5aa735e8-3cbd-11e9-b210-d663bd873d93",
				ExtCreditInstitut: "SEB",
				ExtLoanClearing:   "5270",
				ExtLoanNumber:     "0028600",
				ExtLoanAmount:     100000.00,
				ExtRedeemLoan:     true})
	}
	//
	if err := json.NewEncoder(w).Encode(extloans); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}
