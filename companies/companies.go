package companies

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Company
type company struct {
	ProcessID       string `json:"processId"`
	CompanyID       string `json:"companyId"`
	OrgNumber       string `json:"orgNr"`
	CompanyName     string `json:"companyName"`
	NumberOfLoans   string `json:"numberOfLoans"`
	Created         string `json:"created"`
	BusinessFocus   string `json:"businessFocus"`
	SelectedCompany bool   `json:"selectedCompany"`
}

// companyID
type companyID struct {
	CompanyID string `json:"companyId"`
}

// updateCopmpanyType
type updateCompanyType struct {
	BusinessFocus string `json:"businessFocus"`
}

// CompanyEntry
func CompanyEntry(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	switch r.Method {
	case "POST":
		getCompany(w, r)
	case "DELETE":
		deleteCompany(w, r)
	case "PATCH":
		updateCompany(w, r)
	}
}

// GetCompanies
func GetCompanies(w http.ResponseWriter, r *http.Request) {

	var companies []company
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("getCompanies executed: processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		companies = append(companies,
			company{
				ProcessID:       "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CompanyID:       "461460c2-3d14-11e9-b210-d663bd873d93",
				OrgNumber:       "551010-8474",
				CompanyName:     "Anna Andersson Skog och djurhållning",
				Created:         "2012-01-01",
				SelectedCompany: true})
	}
	//
	if err := json.NewEncoder(w).Encode(companies); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}

//
// getCompany
func getCompany(w http.ResponseWriter, r *http.Request) {

	var companies []company
	processid := r.Header.Get("X-process-Id")
	var data companyID
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
	fmt.Printf("getCompanies executed: processId: %s CompanyId: %s...\r\n", processid, data.CompanyID)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		switch data.CompanyID {
		case "461460c2-3d14-11e9-b210-d663bd873d93":
			companies = append(companies,
				company{
					ProcessID:       "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
					CompanyID:       "461460c2-3d14-11e9-b210-d663bd873d93",
					OrgNumber:       "551010-8474",
					CompanyName:     "Anna Andersson Skog och djurhållning",
					Created:         "2012-01-01",
					SelectedCompany: true})
		}
	}
	//
	if err := json.NewEncoder(w).Encode(companies); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}

// deleteCompany
func deleteCompany(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	//varsh := r.Header
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("Delete-Company executed: processId: %s...\r\n", processid)
	//
	var data companyID
	var r1 []byte
	r1, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewDecoder(bytes.NewReader([]byte(r1))).Decode(&data)
	fmt.Printf("deleteCompany executed: processId: %s/CustomerId: %s...\r\n", processid, data.CompanyID)
	w.WriteHeader(http.StatusOK)
}

// updateCompany
func updateCompany(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	//varsh := r.Header
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("updateCompany executed, processId: %s...\r\n", processid)
	//
	var data updateCompanyType
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
	log.Printf("businessFocus: %s", data.BusinessFocus)
	//
	// We have to write back updateApplicantType
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}
