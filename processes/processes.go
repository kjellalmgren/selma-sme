package processes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//
// Process struct
type Process struct {
	ProcessID          string `json:"processId"`
	CustomerID         []customerID
	ProcessCreatedDate string `json:"processCreatedDate"`
}

// customerID
type customerID struct {
	CustomerID string `json:"customerId"`
}

//
// getProcesses
//
func GetProcesses(w http.ResponseWriter, r *http.Request) {

	var processes []Process
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
	customerid := r.Header.Get("X-customer-Id")
	fmt.Printf("getProcesses executed, customerID: %s...\r\n", customerid)
	//var r1 []byte
	//r1, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	fmt.Fprintf(w, "%s", err)
	//	fmt.Println("Tester\r\n")
	//	w.WriteHeader(http.StatusNotFound)
	//}
	//
	switch customerid {
	case "19640120-3887":
		processes = append(processes,
			Process{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CustomerID: []customerID{
					customerID{
						CustomerID: "19640120-3887",
					},
				},
				ProcessCreatedDate: "2019-03-01"})
		processes = append(processes,
			Process{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
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
