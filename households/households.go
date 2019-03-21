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

	var households []models.Household
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
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		households = append(households,
			models.Household{
				ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				HouseholdMembers: []models.HouseholdMemberType{
					models.HouseholdMemberType{
						CustomerIDs: "19640120-3887",
					},
				},
				HouseholdID:          "4253017a-3d17-11e9-b210-d663bd873d93",
				NumberOfChildsAtHome: 2,
				Childs: []models.ChildType{
					models.ChildType{
						ChildID:         "248485ca-3d9d-11e9-b210-d663bd873d93",
						ChildsAge:       5,
						PartInHousehold: true,
					},
					models.ChildType{
						ChildID:         "eb38da7c-3d9d-11e9-b210-d663bd873d93",
						ChildsAge:       8,
						PartInHousehold: true,
					},
				},
				NumberOfCars:         1,
				ChildMaintenanceCost: 0})
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

	var households []models.Household
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
		processid, data.HosueholdID)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		switch data.HosueholdID {
		case "4253017a-3d17-11e9-b210-d663bd873d93":
			households = append(households,
				models.Household{
					ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
					HouseholdMembers: []models.HouseholdMemberType{
						models.HouseholdMemberType{
							CustomerIDs: "19640120-3887",
						},
					},
					HouseholdID:          "4253017a-3d17-11e9-b210-d663bd873d93",
					NumberOfChildsAtHome: 2,
					Childs: []models.ChildType{
						models.ChildType{
							ChildID:         "248485ca-3d9d-11e9-b210-d663bd873d93",
							ChildsAge:       5,
							PartInHousehold: true,
						},
						models.ChildType{
							ChildID:         "eb38da7c-3d9d-11e9-b210-d663bd873d93",
							ChildsAge:       8,
							PartInHousehold: true,
						},
					},
					NumberOfCars:         1,
					ChildMaintenanceCost: 0})
		}
	}
	//
	if err := json.NewEncoder(w).Encode(households); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}
