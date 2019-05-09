package maintenancecosts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"selmasme/models"
)

// GetMaintenanceCosts documntation
func GetMaintenanceCosts(w http.ResponseWriter, r *http.Request) {

	//var loans []models.Loan
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("GetMaintenanceCosts executed: processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	maintenancecosts := []models.MaintenanceCostType{}
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		file, err := ioutil.ReadFile("json/maintenancecosts.json")
		if err != nil {
			fmt.Fprintf(w, "Error reading maintenancecosts.json - %s", err)
			w.WriteHeader(http.StatusNotFound)
		}
		_ = json.Unmarshal([]byte(file), &maintenancecosts)
	}
	//
	if err := json.NewEncoder(w).Encode(maintenancecosts); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}

// GetMaintenanceCosts
func GetMaintenanceCost(w http.ResponseWriter, r *http.Request) {

	//var loans []models.Loan
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("GetMaintenanceCost executed: processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	w.WriteHeader(http.StatusOK)
}

// AddMaintenanceCost documentation
func AddMaintenanceCost(w http.ResponseWriter, r *http.Request) {

	//var loans []models.Loan
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("AddMaintenanceCost executed: processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	w.WriteHeader(http.StatusAccepted)
}

// DeleteMaintenanceCost documentation
func DeleteMaintenanceCost(w http.ResponseWriter, r *http.Request) {

	//var loans []models.Loan
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("DeleteMaintenanceCost executed: processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	w.WriteHeader(http.StatusOK)
}
