package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// todo
type Process struct {
	ProcessId  string `json:"processId"`
	CustomerId string `json:"customerId,omitempty"`
}

var processes []Process
var process Process

// our main function
func main() {
	process.ProcessId = "1"
	process.CustomerId = "Kjell Almgren"
	processes = append(processes, Process{ProcessId: "1", CustomerId: "19601005-0190"})
	processes = append(processes, Process{ProcessId: "2", CustomerId: "19710101-0120"})
	processes = append(processes, Process{ProcessId: "3", CustomerId: "19720101-0120"})
	processes = append(processes, Process{ProcessId: "4", CustomerId: "19730202-0240"})
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", Index)
	router.HandleFunc("/v1/Process", getProcess).Methods("GET", "OPTIONS")
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
//	getProcess
//
func getProcess(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("getProcess executed...")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	if err := json.NewEncoder(w).Encode(process); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

//
// getProcesses
//
func getProcesses(w http.ResponseWriter, r *http.Request) {

	processes = append(processes, Process{ProcessId: "11", CustomerId: "19601005-0190"})
	fmt.Printf("getProcesses executed...\r\n")
	for i, p := range processes {
		fmt.Println(i, p.ProcessId, p.CustomerId)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	if err := json.NewEncoder(w).Encode(processes); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}
