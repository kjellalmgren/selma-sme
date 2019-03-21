package main

/*
Services: selma-sme
	Author: Kjell Osse Almgren, Tetracon AB
	Date: 2019-03-02
	Description: Service to feed Selma-SME UX, just for test purpose
	Architecture:
	win32: GOOS=windows GOARCH=386 go build -v
	win64: GOOS=windows GOARCH=amd64 go build -v
	arm64: GOOS=linux GOARCH=arm64 go build -v
	arm: GOOS=linux GOARCH=arm go build -v
	exprimental: GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags '-w -s' -a -installsuffix cgo -o selmasme
	expriemntal: CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -tags selmasme -ldflags '-w'
	exprimental: GOOS=linux GOARCH=arm64 go build -a --ldflags 'extldflags "-static"' -tags selmasme -installsuffix selma-sme .
*/

//

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"selmasme/applicants"
	"selmasme/cases"
	"selmasme/collaterals"
	"selmasme/companies"
	"selmasme/companyeconomies"
	"selmasme/extloans"
	"selmasme/households"
	"selmasme/kycinformations"
	"selmasme/loans"
	"selmasme/personaleconomies"
	"selmasme/processes"
	"selmasme/version"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//
//	Processes
//
// Banner text
const (
	// TETRACON banner
	TETRACON = `
_________    __
|__    __|   | |
   |  |  ___ | |_   ____  ___   ___ ___  _ __ 
   |  | / _ \|  _| /  __|/ _ \ / __/ _ \| '_ \
   |  | \ __/| |_  | |  | (_| | (_| (_) | | | | 
   |__| \___| \__| |_|   \__,_|\___\___/|_| |_| 
Server-version: %s Model-version: %s
`
)

//
var (
	srv  bool
	vrsn bool
)

var (
	client *http.Client
	pool   *x509.CertPool
)

// init
func init() {

	// instanciate a new logger
	var log = logrus.New()
	log.Formatter = new(logrus.TextFormatter)
	log.Level = logrus.DebugLevel
	color.Set(color.FgHiGreen)
	fmt.Fprint(os.Stderr, fmt.Sprintf(TETRACON, version.ServerVersion(), version.ModelVersion()))
	color.Unset()
}

//
// our main function
func main() {

	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", Index)
	// processes.go
	router.HandleFunc("/v1/Processes", processes.GetProcesses).Methods("POST", "PATCH", "PUT", "OPTIONS")
	router.HandleFunc("/v1/Process", processes.DeleteProcess).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/v1/getProcess", processes.GetProcessAll).Methods("GET", "OPTIONS")
	// cases.go
	router.HandleFunc("/v1/reserveCaseId", cases.ReserveCaseID).Methods("POST", "GET", "OPTIONS")
	router.HandleFunc("/v1/setCaseIdStatus", cases.SetCaseIDStatus).Methods("PATCH", "OPTIONS")
	// applicants.go
	router.HandleFunc("/v1/Applicants", applicants.GetApplicants).Methods("GET", "POST", "PATCH", "DELETE", "OPTIONS")
	router.HandleFunc("/v1/Applicant", applicants.ApplicantEntry).Methods("GET", "POST", "PATCH", "DELETE", "OPTIONS")
	router.HandleFunc("/v1/Applicant", applicants.ApplicantEntry).Methods("GET", "POST", "PATCH", "DELETE", "OPTIONS")
	router.HandleFunc("/v1/Applicant", applicants.ApplicantEntry).Methods("GET", "POST", "PATCH", "DELETE", "OPTIONS")
	router.HandleFunc("/v1/Applicant", applicants.ApplicantEntry).Methods("PUT", "OPTIONS")
	// loans.go
	router.HandleFunc("/v1/Loans", loans.GetLoans).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/Loan", loans.GetLoan).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/Loan", loans.DeleteLoan).Methods("DELETE", "OPTIONS")
	//
	router.HandleFunc("/v1/xloans", loans.GetLoansx).Methods("POST", "GET", "OPTIONS")
	router.HandleFunc("/v1/xloans", loans.GetLoanx).Methods("POST", "GET", "OPTIONS")
	router.HandleFunc("/v1/xloans", loans.DeleteLoanx).Methods("POST", "DELETE", "OPTIONS")
	// extloans.go
	router.HandleFunc("/v1/ExtLoans", extloans.GetExtLoans).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/ExtLoan", extloans.GetExtLoan).Methods("POST", "OPTIONS")
	// collaterals.go
	router.HandleFunc("/v1/Collaterals", collaterals.GetCollaterals).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/Collateral", collaterals.GetCollateral).Methods("POST", "OPTIONS")
	// companies.go
	router.HandleFunc("/v1/Companies", companies.GetCompanies).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/Company", companies.CompanyEntry).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/Company", companies.CompanyEntry).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/v1/Company", companies.CompanyEntry).Methods("PATCH", "OPTIONS")
	// households.go
	router.HandleFunc("/v1/Households", households.GetHouseholds).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/Household", households.GetHousehold).Methods("POST", "OPTIONS")
	// personalEconomies
	router.HandleFunc("/v1/PersonalEconomies", personaleconomies.GetPersonalEconomies).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/PersonalEconomy", personaleconomies.GetPersonalEconomy).Methods("POST", "OPTIONS")
	// companyEconomies
	router.HandleFunc("/v1/CompanyEconomies", companyeconomies.GetCompanyEconomies).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/CompanyEconomy", companyeconomies.CompanyEconomyEntry).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/CompanyEconomy", companyeconomies.CompanyEconomyEntry).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/v1/CompanyEconomy}", companyeconomies.CompanyEconomyEntry).Methods("PATCH", "OPTIONS")
	// kycinformations.go
	router.HandleFunc("/v1/KycInformations", kycinformations.GetKycInformations).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/KycInformation", kycinformations.GetKycInformation).Methods("POST", "OPTIONS")
	// Healt services local
	router.HandleFunc("/v1/ping", HealthCheckHandler).Methods("GET")
	color.Set(color.FgHiRed)
	fmt.Fprint(os.Stdout, fmt.Sprintf(getHostname()))
	color.Set(color.FgHiGreen)
	fmt.Printf(" Listen on TLS-server ")
	color.Set(color.FgHiYellow)
	fmt.Printf("localhost")
	color.Set(color.FgHiCyan)
	fmt.Printf(":8443\r\n")
	color.Unset()
	err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", router)
	if err != nil {
		fmt.Printf("ListenAndServeTLS Error: %s", err.Error())
		log.Fatal(err)
		logrus.Fatal(err)
	}
}

//
// Index
//
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(os.Stderr, fmt.Sprintf(TETRACON, version.ServerVersion(), version.ModelVersion()))
}

//
//	getHostName
//
func getHostname() string {

	hostname, err1 := os.Hostname()
	if err1 != nil {
		//log.Printf("Hostname: %s", hostname)
		fmt.Println("Error when try to resolve Hostname: ", hostname)
	}
	return hostname
}

//
// HealthCheckHandler
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	//w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	//io.WriteString(w, `[{"alive": "true",`)
	//io.WriteString(w, `"status": "200"}]`)

	type HealthCheck struct {
		HealthStatus int `json:"healthStatus,omitempty"`
	}
	var healtCheck HealthCheck
	healtCheck.HealthStatus = http.StatusOK
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//
	if err := json.NewEncoder(w).Encode(healtCheck); err != nil {
		panic(err)
	}
	fmt.Printf("Http-Status %d received\r\n", http.StatusOK)
}
