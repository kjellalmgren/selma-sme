package takeoverloans

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"selmasme/models"
)

// GetTakeoverLoans method
func GetTakeoverLoans(w http.ResponseWriter, r *http.Request) {

	processid := r.Header.Get("X-process-Id")
	fmt.Printf("getTakeoverLoans executed: processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
}

// GetTakeoverLoans method
func GetTakeoverLoan(w http.ResponseWriter, r *http.Request) {

	processid := r.Header.Get("X-process-Id")
	fmt.Printf("getTakeoverLoan executed: processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")

	var data models.TakeoverLoanID
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
	fmt.Printf("getTakeoverLoan executed: X-process-ID: %s...LoanId: %s\r\n", processid, data.TakeoverLoanID)
	//
	takeoverloans := []models.TakeoverLoanType{}
	file, err := ioutil.ReadFile("json/takeoverloans.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading takeoverloans.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	takeoverloanret := make([]models.TakeoverLoanType, 1, 1)
	_ = json.Unmarshal([]byte(file), &takeoverloans)
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		for i := 0; i < len(takeoverloans); i++ {
			if takeoverloans[i].TakeoverLoanID == data.TakeoverLoanID {
				takeoverloanret[0] = takeoverloans[i]
			}
		}
	}
	//
	if takeoverloanret[0].TakeoverLoanID == data.TakeoverLoanID {
		if err := json.NewEncoder(w).Encode(takeoverloanret); err != nil {
			w.WriteHeader(http.StatusNotFound)
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "TakeoverLoan not Found", 404)
		//w.WriteHeader(http.StatusNotFound)
	}
	//
}
