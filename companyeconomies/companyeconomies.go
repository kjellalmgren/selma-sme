package companyeconomies

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"selmasme/models"
)

//
// GetCompanyEconomies documentation
func GetCompanyEconomies(w http.ResponseWriter, r *http.Request) {

	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("getCompanyEconomies executed: processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	companyeconomies := []models.CompanyEconomy{}
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		file, err := ioutil.ReadFile("json/companyeconomies.json")
		if err != nil {
			fmt.Fprintf(w, "Error reading companyeconomies.json - %s", err)
			w.WriteHeader(http.StatusNotFound)
		}
		_ = json.Unmarshal([]byte(file), &companyeconomies)
	}
	//
	if err := json.NewEncoder(w).Encode(companyeconomies); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}

// CompanyEconomyEntry
func CompanyEconomyEntry(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	switch r.Method {
	case "POST":
		getCompanyEconomy(w, r)
	case "DELETE":
		deleteCompanyEconomy(w, r)
	case "PATCH":
		updateCompanyEconomy(w, r)
	}
}

// getCompanyEconomy documentation
func getCompanyEconomy(w http.ResponseWriter, r *http.Request) {

	//
	processid := r.Header.Get("X-process-Id")
	//
	var data models.CompanyEconomyID
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
	fmt.Printf("getCompanyEconomy executed: processId: %s / companyEconomyId: %s...\r\n",
		processid, data.CompanyEconomyID)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	companyeconomies := []models.CompanyEconomy{}
	file, err := ioutil.ReadFile("json/companyeconomies.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading companyeconomies.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	ceconret := make([]models.CompanyEconomy, 1, 1)
	_ = json.Unmarshal([]byte(file), &companyeconomies)
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		for i := 0; i < len(companyeconomies); i++ {
			if companyeconomies[i].CompanyEconomyID == data.CompanyEconomyID {
				ceconret[0] = companyeconomies[i]
			}
		}
	}
	if ceconret[0].CompanyEconomyID == data.CompanyEconomyID {
		if err := json.NewEncoder(w).Encode(ceconret); err != nil {
			w.WriteHeader(http.StatusNotFound)
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "CompanyEconomy not Found", 404)
		//w.WriteHeader(http.StatusNotFound)
	}
}

func deleteCompanyEconomy(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	//varsh := r.Header
	processid := r.Header.Get("X-process-Id")
	//
	var data models.CompanyEconomyID
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
	fmt.Printf("deleteCompanyEconomy executed, processId: %s companyId: %s...\r\n", processid, data.CompanyEconomyID)
	w.WriteHeader(http.StatusOK)
}

//
// updateCompanyEconomy documentation
func updateCompanyEconomy(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	//varsh := r.Header
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("updateCompanyEconomy executed, processId: %s...\r\n", processid)
	w.WriteHeader(http.StatusOK)
}
