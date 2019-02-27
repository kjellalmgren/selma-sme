package main

//
/*
Services: selma-sme
	Author: Kjell Osse Almgren, Tetracon AB
	Date: 2019-02-27
	Description: Service to feed Selma-SME UX, just for test purpose
	Architecture:
	win32: GOOS=windows GOARCH=386 go build -v
	win64: GOOS=windows GOARCH=amd64 go build -v
	arm64: GOOS=linux GOARCH=arm64 go build -v
	arm: GOOS=linux GOARCH=arm go build -v
	exprimental: GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags '-w -s' -a -installsuffix cgo -o selma-sme
	expriemntal: CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -tags selma-sme -ldflags '-w'
	exprimental: GOOS=linux GOARCH=arm64 go build -a --ldflags 'extldflags "-static"' -tags selma-sme -installsuffix selma-sme .
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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

// **************************************
//
type Process struct {
	ProcessID   string `json:"processId,omitempty"`
	CustomerID  string `json:"customerId,omitempty"`
	DateCreated string `json:"dateCreated,omitempty"`
}

// our main function
func main() {

	fmt.Fprint(os.Stderr, fmt.Sprintf(TETRACON, PingVersion()))
	fmt.Fprint(os.Stdout, fmt.Sprintf(getHostname()))
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", Index)
	router.HandleFunc("/v1/Processes/{customerId}", getProcesses).Methods("GET", "OPTIONS")
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
// getProcesses
//
func getProcesses(w http.ResponseWriter, r *http.Request) {

	var processes []Process
	//
	fmt.Printf("getProcesses executed...\r\n")
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//
	vars := mux.Vars(r)
	fmt.Println(vars["customerId"]) // customerId
	fmt.Println()
	switch vars["customerId"] {
	case "19700101-1001":
		processes = append(processes, Process{ProcessID: "11", CustomerID: "19700101-1001", DateCreated: "2019-02-27"})
	case "19700101-2002":
		processes = append(processes, Process{ProcessID: "22", CustomerID: "19700101-2002", DateCreated: "2019-02-20"})
	default:
		w.WriteHeader(http.StatusNotFound)
	}
	//w.WriteHeader(http.StatusNotFound)

	//
	for i, p := range processes {
		fmt.Println(i, p.ProcessID, p.CustomerID)
	}
	//
	if err := json.NewEncoder(w).Encode(processes); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

//
//	getHostName
//
func getHostname() string {

	hostname, err1 := os.Hostname()
	if err1 == nil {
		//log.Printf("Hostname: %s", hostname)
		fmt.Println("Error when try to resolve Hostname: ", hostname)
	}
	return hostname
}

//
//	getVersion
func PingVersion() string {
	return "v0.4.8"
}
