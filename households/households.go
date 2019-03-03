package households

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Household
type Household struct {
	ProcessID            string `json:"processId,omitempty"`
	HouseholdMembers     []HouseholdMemberType
	HouseholdID          string `json:"householdId,omitempty"`
	NumberOfChildsAtHome string `json:"numberOfChildsAtHome"`
	NumberOfCars         string `json:"numberOfCars"`
	ChildMaintenanceCost string `json:"childMaintenanceCost"`
}

// HouseholdMemberType
type HouseholdMemberType struct {
	CustomerIDs string `json:"householdMember"`
}

// GetApplicants
func GetHouseholds(w http.ResponseWriter, r *http.Request) {

	var households []Household
	vars := mux.Vars(r)
	fmt.Printf("getHouseholds executed: processId: %s...\r\n", vars["processId"])
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//
	switch vars["processId"] {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		households = append(households,
			Household{
				ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				HouseholdMembers: []HouseholdMemberType{
					HouseholdMemberType{
						CustomerIDs: "19640120-3887",
					},
				},
				HouseholdID:          "4253017a-3d17-11e9-b210-d663bd873d93",
				NumberOfChildsAtHome: "2",
				NumberOfCars:         "1",
				ChildMaintenanceCost: "0"})
	}

	//
	if err := json.NewEncoder(w).Encode(households); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}