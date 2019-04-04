package extloans

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"selmasme/models"
)

//
// GetExtLoans
func GetExtLoans(w http.ResponseWriter, r *http.Request) {

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
	extloans := []models.ExtLoanType{}
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		file, err := ioutil.ReadFile("json/extloans.json")
		if err != nil {
			fmt.Fprintf(w, "Error reading extloans.json - %s", err)
			w.WriteHeader(http.StatusNotFound)
		}
		_ = json.Unmarshal([]byte(file), &extloans)
	}
	//
	if err := json.NewEncoder(w).Encode(extloans); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

// GetExtLoan
func GetExtLoan(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	processid := r.Header.Get("X-process-Id")
	var data models.ExtLoanID
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
		data.ExtLoanID)
	//
	extloans := []models.ExtLoanType{}
	file, err := ioutil.ReadFile("json/extloans.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading extloans.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	extloanret := make([]models.ExtLoanType, 1, 1)
	_ = json.Unmarshal([]byte(file), &extloans)
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		for i := 0; i < len(extloans); i++ {
			if extloans[i].ExtLoanID == data.ExtLoanID {
				extloanret[0] = extloans[i]
			}
		}
	}
	if extloanret[0].ExtLoanID == data.ExtLoanID {
		if err := json.NewEncoder(w).Encode(extloanret); err != nil {
			w.WriteHeader(http.StatusNotFound)
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "ExtLoan not Found", 404)
		//w.WriteHeader(http.StatusNotFound)
	}
}
