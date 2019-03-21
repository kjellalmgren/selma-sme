package processes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"selmasme/models"
)

type personalEconomy struct {
	ProcessID          string  `json:"processId"`
	CustomerID         string  `json:"customerId"`
	PersonalEconomyID  string  `json:"personalEconomyId"`
	YearlyIncome       float32 `json:"yearlyIncome"`
	Income             float32 `json:"income"`
	TypeOfEmployeement string  `json:"typeOfEmployeement"`
	Employeer          string  `json:"employeer"`
	EmployeedFromYear  string  `json:"yearOfEmployment"`
}

// CompanyEconomies
type companyEconomy struct {
	ProcessID        string `json:"processId"`
	CompanyID        string `json:"companyId"`
	CompanyEconomyID string `json:"companyEconomyId"`
	Revenues         []revenue
}

type revenue struct {
	RevenueID   string  `json:"revenueId"`
	RevenueYear int32   `json:"revenueYear"`
	Revenue     float32 `json:"revenue"`
}

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

//
// Process struct
type process struct {
	ProcessID          string `json:"processId"`
	CustomerID         []customerID
	ProcessCreatedDate string `json:"processCreatedDate"`
}

// Household
type household struct {
	ProcessID            string `json:"processId"`
	householdMembers     []householdMemberType
	HouseholdID          string `json:"householdId"`
	NumberOfChildsAtHome int    `json:"numberOfChildsAtHome"`
	Childs               []childType
	NumberOfCars         int     `json:"numberOfCars"`
	ChildMaintenanceCost float32 `json:"childMaintenanceCost"`
}

// HouseholdMemberType
type householdMemberType struct {
	CustomerIDs string `json:"householdMember"`
}

//
// ChildType
type childType struct {
	ChildID         string `json:"childId"`
	ChildsAge       int    `json:"childsAge"`
	PartInHousehold bool   `json:"partInHousehold"`
}

// ExtLoan
type extLoan struct {
	ProcessID         string `json:"processId"`
	ExtLoanOwners     []extLoanOwner
	ExtLoanID         string  `json:"extloanId"`
	ExtCreditInstitut string  `json:"extCreditInstitut"`
	ExtLoanClearing   string  `json:"extLoanClearing"`
	ExtLoanNumber     string  `json:"extloanNumber"`
	ExtLoanAmount     float32 `json:"extLoanAmount"`
	ExtRedeemLoan     bool    `json:"extRedeemLoan"`
}

//
// ExtLoanOwner
type extLoanOwner struct {
	CustomerID string `json:"customerId"`
}

// customerID
type customerID struct {
	CustomerID string `json:"customerId"`
}

// processAll
type processAll struct {
	Processes         []process
	Applicans         []applicant
	Loans             []loan
	ExtLoans          []extLoan
	Households        []household
	Companies         []company
	CompanyEconomies  []companyEconomy
	PersonalEconomies []personalEconomy
}

//
// ***************************************************************
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

// ContactInformationType struct
type contactInformationType struct {
	ApplicantEmail        string `json:"applicantByeMail"`
	ApplicantMobileNumber string `json:"applicantBySms"`
}

// Loan
type loan struct {
	ProcessID     string  `json:"processId"`
	LoanID        string  `json:"loanId"`
	LoanNumber    string  `json:"loanNumber"`
	LoanAmount    float32 `json:"loanAmount"`
	PurposeOfLoan string  `json:"purposeOfLoan"`
	Aims          []aimType
}

// aimType
type aimType struct {
	AimID          string  `json:"aimId"`
	AimText        string  `json:"aimText"`
	LoanAmountPart float32 `json:"loanAmountPart"`
}

// ***************************************************************
//
// getProcesses
func GetProcesses(w http.ResponseWriter, r *http.Request) {

	var processes []process
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
	var data customerID
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
	switch data.CustomerID {
	case "19640120-3887":
		processes = append(processes,
			process{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CustomerID: []customerID{
					customerID{
						CustomerID: "19640120-3887",
					},
				},
				ProcessCreatedDate: "2019-03-01"})
		processes = append(processes,
			process{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CustomerID: []customerID{
					customerID{
						CustomerID: "19640120-3887",
					},
				},
				ProcessCreatedDate: "2019-02-26"})
	default:
		//w.WriteHeader(http.StatusNotFound)
	}
	//w.WriteHeader(http.StatusNotFound)

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
	var processes []process
	var applicants []applicant
	var loans []loan
	var extloans []extLoan
	var households []household
	var companies []company
	var companyeconomies []companyEconomy
	var personaleconomies []personalEconomy

	var d models.XAll
	var p models.Person
	p.Name = "Kjell Almgren"
	p.Mobile = "0733981482"
	d.Persons = append(d.Persons, p)

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
		process{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
			CustomerID: []customerID{
				customerID{
					CustomerID: "19640120-3887",
				},
			},
			ProcessCreatedDate: "2019-03-01"})
	//
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
	//
	loans = append(loans,
		loan{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
			LoanID:        "9b8e4822-3cb7-11e9-b210-d663bd873d93",
			LoanNumber:    "930101011212",
			LoanAmount:    2300000,
			PurposeOfLoan: "Köp",
			Aims: []aimType{
				aimType{
					AimID:          "fce3d0aa-4b04-11e9-8646-d663bd873d93",
					AimText:        "Fastighetsköp - annan fastighet",
					LoanAmountPart: 2000000,
				},
				aimType{
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

	//
	if err := json.NewEncoder(w).Encode(processall); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)

}
