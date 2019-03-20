package extloans

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//
// ExtLoan
type extLoan struct {
	ProcessID         string `json:"processId"`
	ExtLoanOwners     []extLoanOwner
	ExtLoanID         string  `json:"extloanId"`
	ExtCreditInstitut string  `json:"extCreditInstitut"`
	ExtLoanClearing   string  `json:"extLoanClearing"`
	ExtLoanNumber     string  `json:"extloanNumber"`
	ExtLoanAmount     float32 `json:"extLoanAmount"`
	ExtRedeemLoan     bool    `json:"extRedeemLoan"`
}

//
// ExtLoanOwner
type extLoanOwner struct {
	CustomerID string `json:"customerId"`
}

// extloanID
type extLoanID struct {
	ExtLoanId string `json:"extloanId"`
}

// GetExtLoans
func GetExtLoans(w http.ResponseWriter, r *http.Request) {

	var extloans []extLoan
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("getExtLoans executed: processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		extloans = append(extloans,
			extLoan{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				ExtLoanOwners: []extLoanOwner{
					extLoanOwner{
						CustomerID: "19640120-3887",
					},
				},
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

// GetExtLoan
func GetExtLoan(w http.ResponseWriter, r *http.Request) {
	var extloans []extLoan
	//
	processid := r.Header.Get("X-process-Id")
	var data extLoanID
	var r1 []byte
	r1, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	//
	json.NewDecoder(bytes.NewReader([]byte(r1))).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Printf("getExtLoans executed: processId: %s /extloanId: %s...\r\n",
		processid,
		data.ExtLoanId)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		switch data.ExtLoanId {
		case "5aa735e8-3cbd-11e9-b210-d663bd873d93":
			extloans = append(extloans,
				extLoan{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
					ExtLoanOwners: []extLoanOwner{
						extLoanOwner{
							CustomerID: "19640120-3887",
						},
					},
					ExtLoanID:         "5aa735e8-3cbd-11e9-b210-d663bd873d93",
					ExtCreditInstitut: "SEB",
					ExtLoanClearing:   "5270",
					ExtLoanNumber:     "0028600",
					ExtLoanAmount:     100000.00,
					ExtRedeemLoan:     true})
		}
	}
	//
	if err := json.NewEncoder(w).Encode(extloans); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}
