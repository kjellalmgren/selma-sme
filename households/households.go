package households

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"selmasme/models"
)

// GetHouseholds
func GetHouseholds(w http.ResponseWriter, r *http.Request) {

	//var households []models.Household
	//vars := mux.Vars(r)
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("getHouseholds executed: processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	households := []models.Household{}
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		file, err := ioutil.ReadFile("json/households.json")
		if err != nil {
			fmt.Fprintf(w, "Error reading households.json - %s", err)
			w.WriteHeader(http.StatusNotFound)
		}
		//households = []models.Household{}
		_ = json.Unmarshal([]byte(file), &households)
	}
	//
	if err := json.NewEncoder(w).Encode(households); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}

// GetHousehold
func GetHousehold(w http.ResponseWriter, r *http.Request) {

	//var households []models.Household
	//vars := mux.Vars(r)
	processid := r.Header.Get("X-process-Id")
	//
	var data models.HouseholdID
	//
	var r1 []byte
	r1, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	//
	json.NewDecoder(bytes.NewReader([]byte(r1))).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	//
	fmt.Printf("getHouseholds executed: processId: %s/ HouseholdId: %s...\r\n",
		processid, data.HouseholdID)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	households := []models.Household{}
	file, err := ioutil.ReadFile("json/households.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading households.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	houseret := make([]models.Household, 1, 1)
	_ = json.Unmarshal([]byte(file), &households)
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		for i := 0; i < len(households); i++ {
			if households[i].HouseholdID == data.HouseholdID {
				houseret[0] = households[i]
			}
		}
	}
	//
	if houseret[0].HouseholdID == data.HouseholdID {
		if err := json.NewEncoder(w).Encode(houseret); err != nil {
			w.WriteHeader(http.StatusNotFound)
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Household not Found", 404)
		//w.WriteHeader(http.StatusNotFound)
	}
	//
}
