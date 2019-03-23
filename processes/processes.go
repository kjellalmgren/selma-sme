package processes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"selmasme/models"
)

//
// processAll
type processAll struct {
	Processes         []models.Process
	Applicans         []models.Applicant
	Loans             []models.Loan
	ExtLoans          []models.ExtLoan
	Households        []models.Household
	Companies         []models.Company
	CompanyEconomies  []models.CompanyEconomy
	PersonalEconomies []models.PersonalEconomy
	Collaterals       []models.Collateral
	KycInformations   []models.KycInformation
}

//
// ***************************************************************
//
// GetProcesses
func GetProcesses(w http.ResponseWriter, r *http.Request) {

	//var processes []models.Process
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID, X-customer-ID")
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("getProcesses executed, processid: %s...\r\n", processid)
	//
	//var data customerID
	//customerid := r.Header.Get("X-customer-Id")
	//fmt.Printf("getProcesses executed, customerID: %s...\r\n", customerid)
	//
	var data models.CustomerID
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
	fmt.Printf("getProcesses executed, customerID: %s...\r\n", data.CustomerID)
	//
	file, err := ioutil.ReadFile("processes.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading processes.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	processes := []models.Process{}
	_ = json.Unmarshal([]byte(file), &processes)
	//
	//
	//for i, p := range processes {
	//	fmt.Println(i, p.ProcessID, p.CustomerID)
	//}
	//
	if err := json.NewEncoder(w).Encode(processes); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

//
// DeleteProcess
func DeleteProcess(w http.ResponseWriter, r *http.Request) {

	processid := r.Header.Get("X-process-Id")
	fmt.Printf("deleteProcess executed, processid: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID, caseIdStatus")
	//
	caseidstatus := r.Header.Get("caseIdStatus")
	fmt.Printf("deleteProcess executed, caseIdStatus: %s...\r\n", caseidstatus)
	w.WriteHeader(http.StatusOK)
	//
}

// GetProcessAll
func GetProcessAll(w http.ResponseWriter, r *http.Request) {
	//
	var processall processAll
	var processes []models.Process
	var applicants []models.Applicant
	var loans []models.Loan
	var extloans []models.ExtLoan
	var households []models.Household
	var companies []models.Company
	var companyeconomies []models.CompanyEconomy
	var personaleconomies []models.PersonalEconomy
	var kycinformations []models.KycInformation
	var collaterals []models.Collateral

	var d models.XAll
	var p models.Person
	p.Name = "Kjell Almgren"
	p.Mobile = "0733981482"
	d.Persons = append(d.Persons, p)
	//
	//
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("GetProcessAll executed, processid: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID, caseIdStatus")
	//
	processes = append(processes,
		models.Process{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
			CustomerID: []models.CustomerID{
				models.CustomerID{
					CustomerID: "19640120-3887",
				},
			},
			ProcessCreatedDate: "2019-03-01"})
	//
	applicants = append(applicants,
		models.Applicant{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
			CustomerID:           "19640120-3887",
			ApplicantID:          "12ab301d-b0ae-46ba-ac99-ff7389fe356e",
			ApplicantName:        "Anna Andersson",
			ApplicantAddress:     "Stora vägen 1",
			ApplicantPostAddress: "420 20 Katrineholm",
			StakeholderType:      "",
			ContactInformation: []models.ContactInformationType{
				models.ContactInformationType{
					ApplicantEmail:        "anna.andersson@gmail.com",
					ApplicantMobileNumber: "07344455666",
				},
			},
			ApplicantEmployeed:    false,
			ApplicantLPEmployment: "PERMANENT",
			ApplicantMember:       false,
			ApplicantBySms:        true,
			ApplicantByeMail:      true})
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
	//
	//
	processall.Applicans = append(applicants)
	processall.Processes = append(processes)
	processall.Loans = append(loans)
	processall.ExtLoans = append(extloans)
	processall.Households = append(households)
	processall.Companies = append(companies)
	processall.CompanyEconomies = append(companyeconomies)
	processall.PersonalEconomies = append(personaleconomies)
	processall.KycInformations = append(kycinformations)
	processall.Collaterals = append(collaterals)

	//
	if err := json.NewEncoder(w).Encode(processall); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)

}
