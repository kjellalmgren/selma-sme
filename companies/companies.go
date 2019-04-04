package companies

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"selmasme/models"
)

// CompanyEntry documentations
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
	case "PUT":
		addCompany(w, r)
	}
}

// GetCompanies documentation
func GetCompanies(w http.ResponseWriter, r *http.Request) {

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
	companies := []models.CompanyType{}
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		file, err := ioutil.ReadFile("json/companies.json")
		if err != nil {
			fmt.Fprintf(w, "Error reading companies.json - %s", err)
			w.WriteHeader(http.StatusNotFound)
		}
		//companies = []models.Company{}
		_ = json.Unmarshal([]byte(file), &companies)
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
// getCompany documentation
func getCompany(w http.ResponseWriter, r *http.Request) {

	//var companies []models.Company
	processid := r.Header.Get("X-process-Id")
	var data models.CompanyID
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
	companies := []models.CompanyType{}
	file, err := ioutil.ReadFile("json/companies.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading companies.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	compret := make([]models.CompanyType, 1, 1)
	_ = json.Unmarshal([]byte(file), &companies)
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":

		for i := 0; i < len(companies); i++ {
			if companies[i].CompanyID == data.CompanyID {
				compret[0] = companies[i]
			}
		}
	}
	//
	if compret[0].CompanyID == data.CompanyID {
		if err := json.NewEncoder(w).Encode(compret); err != nil {
			w.WriteHeader(http.StatusNotFound)
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Company not Found", 404)
		//w.WriteHeader(http.StatusNotFound)
	}
	//
}

// deleteCompany documentation
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
	var data models.CompanyID
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

// updateCompany documentation
func updateCompany(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("updateCompany executed, processId: %s...\r\n", processid)
	//
	var data models.UpdateCompanyType
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

// addCompany documentation
func addCompany(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")

	processid := r.Header.Get("X-process-Id")
	fmt.Printf("Delete-Company executed: processId: %s...\r\n", processid)

	var data models.CompanyType
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
	str := fmt.Sprintf("Company %s has been added", data.CompanyName)
	log.Printf(str)
	//http.Error(w, str, 200)
	w.WriteHeader(http.StatusOK)
}
