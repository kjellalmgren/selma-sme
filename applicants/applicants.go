package applicants

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"selmasme/models"
)

//
// GetApplicants
func GetApplicants(w http.ResponseWriter, r *http.Request) {

	//var applicants []models.Applicant
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("getApplicants executed: processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	file, err := ioutil.ReadFile("applicants.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading applicants.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	applicants := []models.Applicant{}
	_ = json.Unmarshal([]byte(file), &applicants)
	//
	if err := json.NewEncoder(w).Encode(applicants); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}

// ApplicantEntry
func ApplicantEntry(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	switch r.Method {
	case "POST":
		getApplicant(w, r)
	case "DELETE":
		deleteApplicant(w, r)
	case "PATCH":
		updateApplicant(w, r)
	case "PUT":
		addApplicant(w, r)
	}
}

//
// GetApplicant
func getApplicant(w http.ResponseWriter, r *http.Request) {

	//var applicants []models.Applicant

	//vars := mux.Vars(r)
	processid := r.Header.Get("X-process-Id")
	var data models.CustomerID
	var r1 []byte
	r1, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error read r.Body - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewDecoder(bytes.NewReader([]byte(r1))).Decode(&data)
	fmt.Printf("getApplicant executed: processId: %s/CustomerId: %s...\r\n", processid, data.CustomerID)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	file, err := ioutil.ReadFile("applicants.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading applicants.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	applicants := []models.Applicant{}
	//appret[0] := []models.Applicant{}
	appret := make([]models.Applicant, 1, 1)
	_ = json.Unmarshal([]byte(file), &applicants)

	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		for i := 0; i < len(applicants); i++ {
			if applicants[i].CustomerID == data.CustomerID {
				appret[0] = applicants[i]
			}
		}
	}
	//
	if applicants[0].CustomerID == data.CustomerID {
		if err := json.NewEncoder(w).Encode(appret); err != nil {
			w.WriteHeader(http.StatusNotFound)
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Applicant not Found", http.StatusNotFound)
	}

	//
}

// "UpdateApplicant..." go-lint
func updateApplicant(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	//varsh := r.Header
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("Update-Applicant executed, processId: %s...\r\n", processid)
	//
	var data models.UpdateApplicantType
	//var d []contactInformationType
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
	log.Printf("stakeHolder: %s", data.StakeholderType)
	//
	d := make([]models.ContactInformationType, 1, 1)
	d[0].ApplicantEmail = "kjell.almgren@landshypotek.se"
	d[0].ApplicantMobileNumber = "+46733981482"
	data.ContactInformation = d
	for _, ci := range data.ContactInformation {
		log.Printf("eMail: %s - MobileNumber: %s", ci.ApplicantEmail, ci.ApplicantMobileNumber)
	}
	// We have to write back updateApplicantType
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteApplicant
func deleteApplicant(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	switch r.Method {
	case "DELETE":
		processid := r.Header.Get("X-process-Id")
		fmt.Printf("Delete-Applicant executed: processId: %s...\r\n", processid)
		//
		var data models.CustomerID
		var r1 []byte
		r1, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "%s", err)
			w.WriteHeader(http.StatusNotFound)
		}
		json.NewDecoder(bytes.NewReader([]byte(r1))).Decode(&data)
		fmt.Printf("deleteApplicant executed: processId: %s/CustomerId: %s...\r\n", processid, data.CustomerID)
		w.WriteHeader(http.StatusOK)
	}
}

// addApplicant
func addApplicant(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("Add-Applicant executed, processId: %s...\r\n", processid)
	w.WriteHeader(http.StatusCreated)
}
