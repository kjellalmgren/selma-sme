package applicants

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Applicant struct
type applicant struct {
	ProcessID             string `json:"processId"`
	CustomerID            string `json:"customerId"`
	ApplicantID           string `json:"applicantId"`
	ApplicantName         string `json:"applicantName"`
	ApplicantAddress      string `json:"applicantAddress"`
	ApplicantPostAddress  string `json:"applicantPostAddress"`
	StakeholderType       string `json:"stakeholderType"`
	ContactInformation    []contactInformationType
	ApplicantEmployeed    bool   `json:"applicantEmployeed"`
	ApplicantLPEmployment string `json:"applicantLPEmployment"`
	ApplicantMember       bool   `json:"applicantMember"`
	ApplicantBySms        bool   `json:"applicantBySms"`
	ApplicantByeMail      bool   `json:"applicantByeMail"`
}

type upateApplicantType struct {
	ContactInformation []contactInformationType
	StakeholderType    string `json:"stakeholderType"`
	ApplicantBySms     bool   `json:"applicantBySms"`
	ApplicantByeMail   bool   `json:"applicantByeMail"`
}

//
type customerID struct {
	CustomerID string `json:"customerId"`
}

// ContactInformationType struct
type contactInformationType struct {
	ApplicantEmail        string `json:"applicantByeMail"`
	ApplicantMobileNumber string `json:"applicantBySms"`
}

//
// GetApplicants
func getApplicants(w http.ResponseWriter, r *http.Request) {

	var applicants []applicant
	//
	//vars := mux.Vars(r)
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
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		applicants = append(applicants,
			applicant{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CustomerID:           "19640120-3887",
				ApplicantID:          "12ab301d-b0ae-46ba-ac99-ff7389fe356e",
				ApplicantName:        "Anna Andersson",
				ApplicantAddress:     "Stora vägen 1",
				ApplicantPostAddress: "420 20 Katrineholm",
				StakeholderType:      "",
				ContactInformation: []contactInformationType{
					contactInformationType{
						ApplicantEmail:        "anna.andersson@gmail.com",
						ApplicantMobileNumber: "07344455666",
					},
				},
				ApplicantEmployeed:    false,
				ApplicantLPEmployment: "PERMANENT",
				ApplicantMember:       false,
				ApplicantBySms:        true,
				ApplicantByeMail:      true})
		applicants = append(applicants,
			applicant{
				ProcessID:            "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CustomerID:           "19650705-5579",
				ApplicantID:          "12ab301d-b0ae-46ba-ac99-ff7389fe356f",
				ApplicantName:        "Patrik Andersson",
				ApplicantAddress:     "Stora vägen 1",
				ApplicantPostAddress: "420 20 Katrineholm",
				StakeholderType:      "",
				ContactInformation: []contactInformationType{
					contactInformationType{
						ApplicantEmail:        "patrik.andersson@katrineholmrevision.se",
						ApplicantMobileNumber: "07335533777",
					},
				},
				ApplicantEmployeed:    false,
				ApplicantLPEmployment: "PERMANENT",
				ApplicantMember:       false,
				ApplicantBySms:        true,
				ApplicantByeMail:      true})
	}
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
	case "GET":
		getApplicants(w, r)
	case "POST":
		getApplicant(w, r)
	case "DELETE":
		deleteApplicant(w, r)
	case "PATCH":
		updateApplicant(w, r)
	}
}

//
// GetApplicant
func getApplicant(w http.ResponseWriter, r *http.Request) {

	var applicants []applicant

	//vars := mux.Vars(r)
	processid := r.Header.Get("X-process-Id")
	var data customerID
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
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		switch data.CustomerID {
		case "19640120-3887":
			applicants = append(applicants,
				applicant{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
					CustomerID:           "19640120-3887",
					ApplicantID:          "12ab301d-b0ae-46ba-ac99-ff7389fe356e",
					ApplicantName:        "Anna Andersson",
					ApplicantAddress:     "Stora vägen 1",
					ApplicantPostAddress: "420 20 Katrineholm",
					StakeholderType:      "BORGENSMAN",
					ContactInformation: []contactInformationType{
						contactInformationType{
							ApplicantEmail:        "anna.andersson@gmail.com",
							ApplicantMobileNumber: "07344455666",
						},
					},
					ApplicantEmployeed:    false,
					ApplicantLPEmployment: "PERMANENT",
					ApplicantMember:       false,
					ApplicantBySms:        true,
					ApplicantByeMail:      true})
		case "19650705-5579":
			applicants = append(applicants,
				applicant{
					ProcessID:            "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
					CustomerID:           "19650705-5579",
					ApplicantID:          "12ab301d-b0ae-46ba-ac99-ff7389fe356f",
					ApplicantName:        "Patrik Andersson",
					ApplicantAddress:     "Stora vägen 1",
					ApplicantPostAddress: "420 20 Katrineholm",
					StakeholderType:      "EKONOMISKINTRESSEGEMENSKAP",
					ContactInformation: []contactInformationType{
						contactInformationType{
							ApplicantEmail:        "patrik.andersson@katrineholmrevision.se",
							ApplicantMobileNumber: "07335533777",
						},
					},
					ApplicantEmployeed:    false,
					ApplicantLPEmployment: "PERMANENT",
					ApplicantMember:       false,
					ApplicantBySms:        true,
					ApplicantByeMail:      true})
		}
	}
	//
	if err := json.NewEncoder(w).Encode(applicants); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
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
	fmt.Printf("Update-Applicant executed: processId: %s...\r\n", processid)
	//
	var data upateApplicantType

	var r1 []byte
	r1, err := ioutil.ReadAll(r.Body)
	if err == nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewDecoder(bytes.NewReader([]byte(r1))).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	log.Printf("stakeHolder: %s", data.StakeholderType)
	//
	for _, ci := range data.ContactInformation {
		log.Printf("eMail: %s - MobileNumber: %s", ci.ApplicantEmail, ci.ApplicantMobileNumber)
	}
	// We have to write back updateApplicantType
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
	//varsh := r.Header
	switch r.Method {
	case "DELETE":
		processid := r.Header.Get("X-process-Id")
		fmt.Printf("Delete-Applicant executed: processId: %s...\r\n", processid)
		//
		var data customerID
		var r1 []byte
		r1, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "%s", err)
			w.WriteHeader(http.StatusNotFound)
		}
		json.NewDecoder(bytes.NewReader([]byte(r1))).Decode(&data)
		fmt.Printf("deleteApplicant executed: processId: %s/CustomerId: %s...\r\n", processid, data.CustomerID)
	}
}
