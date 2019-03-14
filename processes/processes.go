package processes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//
// Process struct
type Process struct {
	ProcessID   string `json:"processId"`
	CustomerID  string `json:"customerId"`
	DateCreated string `json:"dateCreated"`
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
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	vars := mux.Vars(r)

	switch vars["customerId"] {
	case "19640120-3887":
		processes = append(processes,
			Process{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CustomerID:  "19640120-3887",
				DateCreated: "2019-03-01"})
		processes = append(processes,
			Process{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CustomerID:  "19650705-5579",
				DateCreated: "2019-02-26"})
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
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	vars := mux.Vars(r)
	fmt.Println(vars["processId"]) // processId
	w.WriteHeader(http.StatusOK)
	//
}
