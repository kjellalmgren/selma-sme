package processes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Process struct
type Process struct {
	ProcessID   string `json:"processId,omitempty"`
	CustomerID  string `json:"customerId,omitempty"`
	DateCreated string `json:"dateCreated,omitempty"`
}

//
// getProcesses
//
func GetProcesses(w http.ResponseWriter, r *http.Request) {

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
		processes = append(processes, Process{ProcessID: "11", CustomerID: "19700101-2002", DateCreated: "2019-02-26"})
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
// DeleteProcess
func DeleteProcess(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("deleteProcess executed...\r\n")
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//
	vars := mux.Vars(r)
	fmt.Println(vars["processId"]) // processId
	w.WriteHeader(http.StatusOK)
	//
}
