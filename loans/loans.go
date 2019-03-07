package loans

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Loans
type Loan struct {
	ProcessID     string  `json:"processId"`
	LoanID        string  `json:"loanId"`
	LoanNumber    string  `json:"loanNumber"`
	LoanAmount    float32 `json:"loanAmount"`
	NumberOfLoans int     `json:"numberOfLoans"`
	PurposeOfLoan string  `json:"purposeOfLoan"`
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
				LoanAmount:    2000000,
				NumberOfLoans: 1,
				PurposeOfLoan: "FAF"})
		loans = append(loans,
			Loan{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				LoanID:        "76f0b120-3cb8-11e9-b210-d663bd873d93",
				LoanNumber:    "930101011213",
				LoanAmount:    200000,
				NumberOfLoans: 1,
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

//
// GetLoan
func GetLoan(w http.ResponseWriter, r *http.Request) {

	var loans []Loan
	//
	vars := mux.Vars(r)
	fmt.Printf("getLoan executed: processId: %s/customerId: %s...\r\n", vars["processId"], vars["customerId"])
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//
	switch vars["processId"] {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		switch vars["loanId"] {
		case "9b8e4822-3cb7-11e9-b210-d663bd873d93":
			loans = append(loans,
				Loan{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
					LoanID:        "9b8e4822-3cb7-11e9-b210-d663bd873d93",
					LoanNumber:    "930101011212",
					LoanAmount:    2000000,
					NumberOfLoans: 1,
					PurposeOfLoan: "FAF"})
		case "76f0b120-3cb8-11e9-b210-d663bd873d93":
			loans = append(loans,
				Loan{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
					LoanID:        "76f0b120-3cb8-11e9-b210-d663bd873d93",
					LoanNumber:    "930101011213",
					LoanAmount:    200000,
					NumberOfLoans: 1,
					PurposeOfLoan: "RTM"})
		}
	}
	//
	if err := json.NewEncoder(w).Encode(loans); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}

// GetLoansx
func GetLoansx(w http.ResponseWriter, r *http.Request) {

	var loans []Loan
	//
	vars := mux.Vars(r)
	fmt.Printf("XLoans executed: processId: %s...\r\n", vars["processId"])
	fmt.Printf("GetLoansx\n\r")
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	//varsh := r.Header.Get("X-process-ID")
	//w.Header().Set("XprocessID", "9a65d28a-46bb-4442-b96d-6a09fda6b18b")
	// r.Header().Get("XprocessID")
	loans = append(loans,
		Loan{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
			LoanID:        "9b8e4822-3cb7-11e9-b210-d663bd873d93",
			LoanNumber:    "930101011212",
			LoanAmount:    2000000,
			NumberOfLoans: 1,
			PurposeOfLoan: "FAF"})
	loans = append(loans,
		Loan{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
			LoanID:        "76f0b120-3cb8-11e9-b210-d663bd873d93",
			LoanNumber:    "930101011213",
			LoanAmount:    200000,
			NumberOfLoans: 1,
			PurposeOfLoan: "RTM"})
	//
	if err := json.NewEncoder(w).Encode(loans); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
	//
}

// GetLoanx
func GetLoanx(w http.ResponseWriter, r *http.Request) {

	var loans []Loan
	//
	vars := mux.Vars(r)
	fmt.Printf("getLoanx executed: processId: %s/customerId: %s...\r\n", vars["processId"], vars["customerId"])
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	loans = append(loans,
		Loan{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
			LoanID:        "9b8e4822-3cb7-11e9-b210-d663bd873d93",
			LoanNumber:    "930101011212",
			LoanAmount:    2000000,
			NumberOfLoans: 1,
			PurposeOfLoan: "FAF"})

	loans = append(loans,
		Loan{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
			LoanID:        "76f0b120-3cb8-11e9-b210-d663bd873d93",
			LoanNumber:    "930101011213",
			LoanAmount:    200000,
			NumberOfLoans: 1,
			PurposeOfLoan: "RTM"})

	//
	if err := json.NewEncoder(w).Encode(loans); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}

// DeleteLoanx
func DeleteLoanx(w http.ResponseWriter, r *http.Request) {

	//var loans []Loan
	//
	vars := mux.Vars(r)
	fmt.Printf("DeleteLoan executed: processId: %s/customerId: %s...\r\n", vars["processId"], vars["customerId"])
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")

	//
	w.WriteHeader(http.StatusOK)
	//
}
