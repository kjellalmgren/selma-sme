package maintenancecosts

import (
	"bytes"
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

// GetMaintenanceCost documentation
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
	var data models.MaintenanceCostID
	var r1 []byte
	r1, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error read r.Body - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewDecoder(bytes.NewReader([]byte(r1))).Decode(&data)
	fmt.Printf("getMaintenanceCostID executed: processId: %s /MaintenanceCostID: %s...\r\n", processid, data.MaintenanceCostID)

	file, err := ioutil.ReadFile("json/maintenancecosts.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading maintenancecosts.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	maintenancecosts := []models.MaintenanceCostType{}
	//mainret[0] := []models.MaintenanceCostType{}
	mainret := make([]models.MaintenanceCostType, 1, 1)
	_ = json.Unmarshal([]byte(file), &maintenancecosts)
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		for i := 0; i < len(maintenancecosts); i++ {
			//maintenancecosts := []models.MaintenanceCostType{}
			if maintenancecosts[i].MaintenanceCostID == data.MaintenanceCostID {
				mainret[0] = maintenancecosts[i]
			}
		}
	}
	//
	if mainret[0].MaintenanceCostID == data.MaintenanceCostID {
		if err := json.NewEncoder(w).Encode(mainret); err != nil {
			w.WriteHeader(http.StatusNotFound)
			//panic(err)
		}
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Mainrenancecost not Found", http.StatusNotFound)
	}
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
