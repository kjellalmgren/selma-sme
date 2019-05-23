package main

/*
Services: selma-sme
	Author: Kjell Osse Almgren, Tetracon AB
	Date: 2019-05-23
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
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"selmasme/applicants"
	"selmasme/budgets"
	"selmasme/cases"
	"selmasme/collaterals"
	"selmasme/eusupports"
	"selmasme/guarantors"
	"selmasme/maintenancecosts"
	submitapplication "selmasme/submitapplications"
	"selmasme/swagger"

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
Server-version: %s Model-version: %s Model-date: %s
`
)

//
var (
	srv  bool
	vrsn bool
	date bool
)

var (
	client *http.Client
	pool   *x509.CertPool
)

// init documwentation
func init() {

	// instanciate a new logger
	var log = logrus.New()
	log.Formatter = new(logrus.TextFormatter)
	log.Level = logrus.DebugLevel
	color.Set(color.FgHiGreen)
	fmt.Fprint(os.Stderr, fmt.Sprintf(TETRACON, version.ServerVersion(), version.ModelVersion(), version.ModelDate()))
	color.Unset()
}

//
// our main function
func main() {

	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", Index)
	router.HandleFunc("/v1/swagger", swagger.Swagger).Methods("GET", "POST", "OPTIONS")
	router.HandleFunc("/v1/upload", Upload).Methods("GET", "POST", "OPTIONS")
	//
	// processes.go
	router.HandleFunc("/v1/getprocesses", processes.GetProcesses).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/getprocess", processes.GetProcessAll).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/saveprocess", processes.SaveProcessAll).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/deleteprocess", processes.DeleteProcess).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/v1/updateprocessstatus", processes.UpdateProcess).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/v1/addprocess", processes.AddProcess).Methods("PUT", "OPTIONS") // addProcess
	//
	// cases.go
	router.HandleFunc("/v1/reserveCaseId", cases.ReserveCaseID).Methods("POST", "GET", "OPTIONS")
	router.HandleFunc("/v1/setCaseIdStatus", cases.SetCaseIDStatus).Methods("PATCH", "OPTIONS")
	//
	// applicants.go
	router.HandleFunc("/v1/getapplicants", applicants.GetApplicants).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/getapplicant", applicants.GetApplicant).Methods("POST", "GET", "OPTIONS")
	router.HandleFunc("/v1/updateapplicant", applicants.UpdateApplicant).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/v1/deleteapplicant", applicants.DeleteApplicant).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/v1/addapplicant", applicants.AddApplicant).Methods("PUT", "OPTIONS")
	//
	// loans.go
	router.HandleFunc("/v1/getloans", loans.GetLoans).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/getloan", loans.GetLoan).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/deleteloan", loans.DeleteLoan).Methods("DELETE", "OPTIONS")
	//
	// extloans.go
	router.HandleFunc("/v1/getextloans", extloans.GetExtLoans).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/getextloan", extloans.GetExtLoan).Methods("POST", "OPTIONS")
	//
	// collaterals.go
	router.HandleFunc("/v1/getCollaterals", collaterals.GetCollaterals).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/getCollateral", collaterals.GetCollateral).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/updateCollateral", collaterals.UpdateCollateral).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/v1/deleteCollateral", collaterals.DeleteCollateral).Methods("DELETE", "OPTIONS")
	//
	// companies.go
	router.HandleFunc("/v1/getcompanies", companies.GetCompanies).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/getcompany", companies.GetCompany).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/deletecompany", companies.DeleteCompany).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/v1/updatecompany", companies.UpdateCompany).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/v1/addcompany", companies.AddCompany).Methods("PUT", "OPTIONS")
	//
	// Budgets.go
	router.HandleFunc("/v1/getbudgets", budgets.GetBudgets).Methods("POST", "GET", "OPTIONS")
	router.HandleFunc("/v1/getbudget", budgets.GetBudget).Methods("POST", "GET", "OPTIONS")
	//
	// households.go
	router.HandleFunc("/v1/gethouseholds", households.GetHouseholds).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/gethousehold", households.GetHousehold).Methods("POST", "OPTIONS")
	//
	// personalEconomies
	router.HandleFunc("/v1/getpersonaleconomies", personaleconomies.GetPersonalEconomies).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/getpersonaleconomy", personaleconomies.GetPersonalEconomy).Methods("POST", "OPTIONS")
	//
	// companyEconomies
	router.HandleFunc("/v1/gompanyeconomies", companyeconomies.GetCompanyEconomies).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/getcompanyeconomy", companyeconomies.GetCompanyEconomy).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/deletecompanyeconomy", companyeconomies.DeleteCompanyEconomy).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/v1/updatecompanyeconomy}", companyeconomies.UpdateCompanyEconomy).Methods("PATCH", "OPTIONS")
	//
	// kycinformations.go
	router.HandleFunc("/v1/getkycinformations", kycinformations.GetKycInformations).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/getkycinformation", kycinformations.GetKycInformation).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/addkycinformation", kycinformations.AddKycInformation).Methods("PUT", "OPTIONS")
	router.HandleFunc("/v1/deletekycinformation", kycinformations.DeleteKycInformation).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/v1/updatekycinformation", kycinformations.UpdateKycInformation).Methods("PATCH", "OPTIONS")
	//
	// xloans.go
	router.HandleFunc("/v1/getxloans", loans.GetLoansx).Methods("POST", "GET", "OPTIONS")
	router.HandleFunc("/v1/getxloan", loans.GetLoanx).Methods("POST", "GET", "OPTIONS")
	router.HandleFunc("/v1/deletexloan", loans.DeleteLoanx).Methods("DELETE", "OPTIONS")
	//
	// eusupports.go
	router.HandleFunc("/v1/geteusupports", eusupports.GetEUSupports).Methods("POST", "GET", "OPTIONS")
	router.HandleFunc("/v1/geteusupport", eusupports.GetEUSupport).Methods("POST", "GET", "OPTIONS")
	router.HandleFunc("/v1/deleteeusupport", eusupports.DeleteEUSupport).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/v1/addeusupport", eusupports.AddEUSupport).Methods("PUT", "OPTIONS")
	//
	// guarantors.go
	router.HandleFunc("/v1/getguarantors", guarantors.GetGuarantors).Methods("POST", "GET", "OPTIONS")
	router.HandleFunc("/v1/getguarantor", guarantors.GetGuarantor).Methods("POST", "GET", "OPTIONS")
	router.HandleFunc("/v1/addguarantor", guarantors.AddGuarantor).Methods("PUT", "OPTIONS")
	router.HandleFunc("/v1/deleteguarantor", guarantors.DeleteGuarantor).Methods("DELETE", "OPTIONS")
	//
	router.HandleFunc("/v1/getmaintenancecosts", maintenancecosts.GetMaintenanceCosts).Methods("POST", "GET", "OPTIONS")
	router.HandleFunc("/v1/getmaintenancecost", maintenancecosts.GetMaintenanceCost).Methods("POST", "GET", "OPTIONS")
	router.HandleFunc("/v1/addmaintenancecost", maintenancecosts.AddMaintenanceCost).Methods("PUT", "OPTIONS")
	router.HandleFunc("/v1/deletemaintenancecost", maintenancecosts.DeleteMaintenanceCost).Methods("DELETE", "OPTIONS")
	//
	// submitApplication.go
	router.HandleFunc("/v1/submitapplication", submitapplication.Submit).Methods("POST", "OPTIONS")
	// roaring.go
	//router.HandleFunc("/v1/getaccesstoken", roaring.GetAccesToken).Methods("POST", "OPTIONS")
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
	//err := http.ListenAndServe(":8400", router)
	//server := &http.Server{
	//	Addr:         ":8443",
	//	ReadTimeout:  time.Second * 30,
	//	WriteTimeout: time.Second * 30,
	//	IdleTimeout:  time.Second * 30,
	//}
	//var header Header
	// ***
	// header := roaring.GetRoaringAccessToken()
	// fmt.Println(fmt.Sprintf("1:%v", string(header.AccessToken)))
	// fmt.Println(fmt.Sprintf("2:%s", string(header.TokenType)))
	// fmt.Println(string(header.TokenType))
	// ***

	err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", router)
	//err := server.ListenAndServeTLS("cert.pem", "key.pem", router)
	if err != nil {
		fmt.Printf("ListenAndServeTLS Error: %s", err.Error())
		log.Fatal(err)
		logrus.Fatal(err)
	}
}

//
// Index documentation
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(os.Stderr, fmt.Sprintf(TETRACON, version.ServerVersion(), version.ModelVersion()))
}

// Upload documentation
func Upload(w http.ResponseWriter, r *http.Request) {
	//processid := r.Header.Get("X-process-Id")
	//fmt.Printf("swagger executed, processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	processid := r.Header.Get("X-process-Id")
	if processid == "" {
		processid = "9a65d28a-46bb-4442-b96d-6a09fda6b18b"
	}
	fmt.Fprintf(w, "Upload file\n")
	// 1. parse input, type multipart/form-data
	r.ParseMultipartForm(10 << 20)
	// 2. retrieve file from posted form-data
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error retrieving file from form-data")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fn := fmt.Sprintf("%s-%v", processid, handler.Filename)
	fmt.Printf("Uploaded file: %+v\n", handler.Filename)
	fmt.Printf("File Size file: %+v\n", handler.Size)
	fmt.Printf("MIME Header %+v\n", handler.Header)
	// 3.
	fmt.Printf("New filename: %s", fn)
	tempFile, err := ioutil.TempFile("temp-folders/", fmt.Sprintf("%s", fn))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()
	// 4.
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	fmt.Fprintf(w, "Successfully uploaded file \n")

	// Parse
	w.WriteHeader(http.StatusOK)
}

//
//	getHostName documentation
func getHostname() string {

	hostname, err1 := os.Hostname()
	if err1 != nil {
		//log.Printf("Hostname: %s", hostname)
		fmt.Println("Error when try to resolve Hostname: ", hostname)
	}
	return hostname
}

//
// HealthCheckHandler documentation
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
