package applicants

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Applicant struct
type Applicant struct {
	ProcessID             string `json:"processId,omitempty"`
	CustomerID            string `json:"customerId,omitempty"`
	ApplicantID           string `json:"applicantId,omitempty"`
	ApplicantName         string `json:"applicantName,omitempty"`
	ApplicantAddress      string `json:"applicantAddress,omitempty"`
	ApplicantPostAddress  string `json:"applicantPostAddress,omitempty"`
	ApplicantRole         string `json:"applicantRole,omitempty"`
	ContactInformation    []ContactInformationType
	ApplicantEmployeed    bool   `json:"applicantEmployeed,omitempty"`
	ApplicantLPEmployment string `json:"applicantLPEmployment,omitempty"`
	ApplicantMember       bool   `json:"applicantMember"`
	ApplicantBySms        bool   `json:"applicantBySms"`
	ApplicantByeMail      bool   `json:"applicantByeMail"`
}

// ContactInformationType struct
type ContactInformationType struct {
	ApplicantEmail        string `json:"applicantEmail,omitempty"`
	ApplicantMobileNumber string `json:"applicantMobileNumber,omitempty"`
}

//
// GetApplicants
func GetApplicants(w http.ResponseWriter, r *http.Request) {

	var applicants []Applicant
	//
	vars := mux.Vars(r)
	fmt.Printf("getApplicants executed: processId: %s...\r\n", vars["processId"])
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	switch vars["processId"] {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		applicants = append(applicants,
			Applicant{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CustomerID:           "19640120-3887",
				ApplicantID:          "12ab301d-b0ae-46ba-ac99-ff7389fe356e",
				ApplicantName:        "Anna Andersson",
				ApplicantAddress:     "Stora vägen 1",
				ApplicantPostAddress: "420 20 Katrineholm",
				ApplicantRole:        "",
				ContactInformation: []ContactInformationType{
					ContactInformationType{
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
			Applicant{
				ProcessID:            "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CustomerID:           "19650705-5579",
				ApplicantID:          "12ab301d-b0ae-46ba-ac99-ff7389fe356f",
				ApplicantName:        "Patrik Andersson",
				ApplicantAddress:     "Stora vägen 1",
				ApplicantPostAddress: "420 20 Katrineholm",
				ApplicantRole:        "",
				ContactInformation: []ContactInformationType{
					ContactInformationType{
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

//
// GetApplicant
func GetApplicant(w http.ResponseWriter, r *http.Request) {

	var applicants []Applicant

	vars := mux.Vars(r)
	fmt.Printf("getApplicant executed: processId: %s/%s...\r\n", vars["processId"], vars["customerId"])
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	switch vars["processId"] {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		switch vars["customerId"] {
		case "19640120-3887":
			applicants = append(applicants,
				Applicant{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
					CustomerID:           "19640120-3887",
					ApplicantID:          "12ab301d-b0ae-46ba-ac99-ff7389fe356e",
					ApplicantName:        "Anna Andersson",
					ApplicantAddress:     "Stora vägen 1",
					ApplicantPostAddress: "420 20 Katrineholm",
					ApplicantRole:        "",
					ContactInformation: []ContactInformationType{
						ContactInformationType{
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
				Applicant{
					ProcessID:            "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
					CustomerID:           "19650705-5579",
					ApplicantID:          "12ab301d-b0ae-46ba-ac99-ff7389fe356f",
					ApplicantName:        "Patrik Andersson",
					ApplicantAddress:     "Stora vägen 1",
					ApplicantPostAddress: "420 20 Katrineholm",
					ApplicantRole:        "",
					ContactInformation: []ContactInformationType{
						ContactInformationType{
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
func UpdateApplicant(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	//varsh := r.Header
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("Update-Applicant executed: processId: %s...\r\n", processid)
	//
	var data Applicant

	//reader, _ := r.GetBody()
	//defer reader.Close()
	var r1 []byte
	r1, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewDecoder(bytes.NewReader([]byte(r1))).Decode(&data)
	log.Printf("Name: %s", data.ApplicantName)
	w.WriteHeader(http.StatusOK)
}
