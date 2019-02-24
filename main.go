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
	processId  string `json:"processId"`
	customerId string `json:"customerId,omitempty"`
}

var processes []Process
var process Process

// our main function
func main() {
	process.processId = "1"
	process.customerId = "Kjell Almgren"
	processes = append(processes, Process{processId: "1", customerId: "19601005-0190"})
	processes = append(processes, Process{processId: "2", customerId: "19710101-0120"})
	processes = append(processes, Process{processId: "3", customerId: "19720101-0120"})
	processes = append(processes, Process{processId: "4", customerId: "19730202-0240"})
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", Index)
	router.HandleFunc("/v1/getProcess", getProcess).Methods("GET")
	router.HandleFunc("/v1/getProcesses", getProcesses).Methods("GET")
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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	if err := json.NewEncoder(w).Encode(process); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

//
// getProcesses
//
func getProcesses(w http.ResponseWriter, r *http.Request) {

	processes = append(processes, Process{processId: "11", customerId: "19601005-0190"})
	fmt.Printf("getProcesses executed...\r\n")
	for i, p := range processes {
		fmt.Println(i, p.processId, p.customerId)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	if err := json.NewEncoder(w).Encode(processes); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}
