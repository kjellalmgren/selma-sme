package loans

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"selmasme/models"
)

// GetLoans
func GetLoans(w http.ResponseWriter, r *http.Request) {

	var loans []models.Loan
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("getLoans executed: processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		loans = append(loans,
			models.Loan{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				LoanID:        "9b8e4822-3cb7-11e9-b210-d663bd873d93",
				LoanNumber:    "930101011212",
				LoanAmount:    2300000,
				PurposeOfLoan: "Köp",
				Aims: []models.AimType{
					models.AimType{
						AimID:          "fce3d0aa-4b04-11e9-8646-d663bd873d93",
						AimText:        "Fastighetsköp - annan fastighet",
						LoanAmountPart: 2000000,
					},
					models.AimType{
						AimID:          "fce3d0aa-4b04-11e9-8646-d663bd873d94",
						AimText:        "Renovering mjölkstall",
						LoanAmountPart: 300000,
					},
				},
			})
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

	var loans []models.Loan
	//
	processid := r.Header.Get("X-process-Id")
	//
	var data models.LoanID
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
	fmt.Printf("getLoan executed: X-process-ID: %s...LoanId: %s\r\n", processid, data.LoanID)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		switch data.LoanID {
		case "9b8e4822-3cb7-11e9-b210-d663bd873d93":
			loans = append(loans,
				models.Loan{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
					LoanID:        "9b8e4822-3cb7-11e9-b210-d663bd873d93",
					LoanNumber:    "930101011212",
					LoanAmount:    2300000,
					PurposeOfLoan: "Köp",
					Aims: []models.AimType{
						models.AimType{
							AimID:          "fce3d0aa-4b04-11e9-8646-d663bd873d93",
							AimText:        "Fastighetsköp - annan fastighet",
							LoanAmountPart: 2000000,
						},
						models.AimType{
							AimID:          "fce3d0aa-4b04-11e9-8646-d663bd873d94",
							AimText:        "Renovering mjölkstall",
							LoanAmountPart: 300000,
						},
					},
				})
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

	var loans []models.Loan
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("GetXLoans executed: X-process-ID: %s...\r\n", processid)
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
		models.Loan{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
			LoanID:        "9b8e4822-3cb7-11e9-b210-d663bd873d93",
			LoanNumber:    "930101011212",
			LoanAmount:    2300000,
			PurposeOfLoan: "Köp",
			Aims: []models.AimType{
				models.AimType{
					AimID:          "fce3d0aa-4b04-11e9-8646-d663bd873d93",
					AimText:        "Fastighetsköp - annan fastighet",
					LoanAmountPart: 2000000,
				},
				models.AimType{
					AimID:          "fce3d0aa-4b04-11e9-8646-d663bd873d94",
					AimText:        "Renovering mjölkstall",
					LoanAmountPart: 300000,
				},
			},
		})
	//
	if err := json.NewEncoder(w).Encode(loans); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}

// GetLoanx
func GetLoanx(w http.ResponseWriter, r *http.Request) {

	var loans []models.Loan
	processid := r.Header.Get("X-process-Id")
	var data models.LoanID
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
	//
	fmt.Printf("getLoanx executed: X-process-ID: %s/loanId: %s...\r\n", processid, data.LoanID)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	loans = append(loans,
		models.Loan{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
			LoanID:        "9b8e4822-3cb7-11e9-b210-d663bd873d93",
			LoanNumber:    "930101011212",
			LoanAmount:    2300000,
			PurposeOfLoan: "Köp",
			Aims: []models.AimType{
				models.AimType{
					AimID:          "fce3d0aa-4b04-11e9-8646-d663bd873d93",
					AimText:        "Fastighetsköp - annan fastighet",
					LoanAmountPart: 2000000,
				},
				models.AimType{
					AimID:          "fce3d0aa-4b04-11e9-8646-d663bd873d94",
					AimText:        "Renovering mjölkstall",
					LoanAmountPart: 300000,
				},
			},
		})
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

	//
	processid := r.Header.Get("X-process-Id")
	var data models.LoanID
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
	//
	fmt.Printf("DeleteLoanx executed: processId: %s/loanID: %s...\r\n", processid, data.LoanID)
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

// DeleteLoan
func DeleteLoan(w http.ResponseWriter, r *http.Request) {

	//
	processid := r.Header.Get("X-process-Id")
	//
	var data models.LoanID
	//
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
	//
	fmt.Printf("DeleteLoan executed: processId: %s/loanID: %s...\r\n", processid, data.LoanID)
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
