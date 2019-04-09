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
// GetApplicants documentation
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
	file, err := ioutil.ReadFile("json/applicants.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading applicants.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	applicants := []models.ApplicantType{}
	_ = json.Unmarshal([]byte(file), &applicants)
	//
	if err := json.NewEncoder(w).Encode(applicants); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	//
}

// ApplicantEntry documentation, not exposed
func applicantEntry(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	switch r.Method {
	case "POST":
		GetApplicant(w, r)
	case "DELETE":
		DeleteApplicant(w, r)
	case "PATCH":
		UpdateApplicant(w, r)
	case "PUT":
		AddApplicant(w, r)
	}
}

//
// GetApplicant documentation
func GetApplicant(w http.ResponseWriter, r *http.Request) {

	//
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
	file, err := ioutil.ReadFile("json/applicants.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading applicants.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	applicants := []models.ApplicantType{}
	//appret[0] := []models.Applicant{}
	appret := make([]models.ApplicantType, 1, 1)
	_ = json.Unmarshal([]byte(file), &applicants)
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		for i := 0; i < len(applicants); i++ {
			if applicants[i].CustomerID == data.CustomerID {
				appret[0] = applicants[i]
			}
		}
	}
	//
	if appret[0].CustomerID == data.CustomerID {
		if err := json.NewEncoder(w).Encode(appret); err != nil {
			w.WriteHeader(http.StatusNotFound)
			//panic(err)
		}
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Applicant not Found", http.StatusNotFound)
	}
	//
}

// "UpdateApplicant..." go-lint
func UpdateApplicant(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "PATCH, OPTIONS")
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

// DeleteApplicant documentation
func DeleteApplicant(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
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

// addApplicant documentation
// When we create an applicant we should be returning a uuid and start communicate this
// in the API calls going forward.
func AddApplicant(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	// 19640120-3887 d5744655-b71e-428a-98b9-2b6c66c8c95a
	// 19650705-5579 b2f86b36-7ff3-428e-ab45-8dad11952dae
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("Add-Applicant executed, processId: %s...\r\n", processid)
	//
	data := models.ApplicantType{}
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
	if data.CustomerID == "19640122-3887" {
		data.ApplicantID = "d5744655-b71e-428a-98b9-2b6c66c8c95a"
	} else if data.CustomerID == "19650705-5579" {
		data.ApplicantID = "b2f86b36-7ff3-428e-ab45-8dad11952dae"
	}
	applicant := make([]models.ApplicantType, 1, 1)
	applicants := make([]models.ApplicantType, 1, 1)
	applicant[0] = data
	applicants[0] = applicant[0]

	if err := json.NewEncoder(w).Encode(applicants); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	fmt.Printf("CustomerId: %s - Generated ApplicantID: %s",
		data.CustomerID, data.ApplicantID)
	w.WriteHeader(http.StatusCreated)
}
