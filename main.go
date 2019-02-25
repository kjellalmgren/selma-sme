package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//
//	Process
//
type Process struct {
	ProcessID  string `json:"processId"`
	CustomerID string `json:"customerId,omitempty"`
}

//var processes []Process
//var process Process

// our main function
func main() {
	//process.ProcessID = "1"
	//process.CustomerID = "Kjell Almgren"
	//processes = append(processes, Process{ProcessID: "1", CustomerID: "19601005-0190"})
	//processes = append(processes, Process{ProcessID: "2", CustomerID: "19710101-0120"})
	//processes = append(processes, Process{ProcessID: "3", CustomerID: "19720101-0120"})
	//processes = append(processes, Process{ProcessID: "4", CustomerID: "19730202-0240"})
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
// getProcess
//
func getProcess(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("getProcess executed...")

	var process Process

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

	var processes []Process
	//var process Process
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
		processes = append(processes, Process{ProcessID: "11", CustomerID: "19700101-1001"})
	case "19700101-2002":
		processes = append(processes, Process{ProcessID: "22", CustomerID: "19700101-2002"})
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
