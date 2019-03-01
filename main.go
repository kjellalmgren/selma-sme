package main

/*
Services: selma-sme
	Author: Kjell Osse Almgren, Tetracon AB
	Date: 2019-02-28
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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"selmasme/applicants"
	"selmasme/cases"
	"selmasme/processes"
	"selmasme/version"

	"github.com/gorilla/mux"
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
version: %s
`
)

//
var (
	srv  bool
	vrsn bool
)

//
// our main function
func main() {

	fmt.Fprint(os.Stderr, fmt.Sprintf(TETRACON, version.PingVersion()))
	fmt.Fprint(os.Stdout, fmt.Sprintf(getHostname()))
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", Index)
	// processes.go
	router.HandleFunc("/v1/Processes/{customerId}", processes.GetProcesses).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/deleteProcess/{processId}", processes.DeleteProcess).Methods("POST", "OPTIONS")
	// cases.go
	router.HandleFunc("/v1/reserveCaseId/{processId}/{customerId}/{caseIdStatus}", cases.ReserveCaseId).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/setCaseIdStatus/{processId}/{caseId}/{caseIdStatus}", cases.SetCaseIDStatus).Methods("POST", "OPTIONS")
	// applicants.go
	router.HandleFunc("/v1/Applicants/{processId}", applicants.GetApplicants).Methods("GET", "OPTIONS")
	// Healt services local
	router.HandleFunc("/v1/ping", HealthCheckHandler).Methods("GET")
	fmt.Printf("Listen on server localhost:8000\r\n")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Printf("ListenAndServer Error: %s", err.Error())
		log.Fatal(err)
	}
}

//
// Index
//
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
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
//
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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//
	if err := json.NewEncoder(w).Encode(healtCheck); err != nil {
		panic(err)
	}
	fmt.Printf("Http-Status %d received\r\n", http.StatusOK)
}
