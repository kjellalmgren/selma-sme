package personaleconomies

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// PersonalEconomy
type personalEconomy struct {
	ProcessID          string  `json:"processId"`
	CustomerID         string  `json:"customerId"`
	PersonalEconomyID  string  `json:"personalEconomyId"`
	YearlyIncome       float32 `json:"yearlyIncome"`
	Income             float32 `json:"income"`
	TypeOfEmployeement string  `json:"typeOfEmployeement"`
	Employeer          string  `json:"employeer"`
	EmployeedFromYear  string  `json:"yearOfEmployment"`
}

//
type personalEconomyID struct {
	PersonalEconomyID string `json:"personalEconomyID"`
}

// GetPersonalEconomies
func GetPersonalEconomies(w http.ResponseWriter, r *http.Request) {

	var personaleconomies []personalEconomy
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("getPersonalEconomies executed: processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		personaleconomies = append(personaleconomies,
			personalEconomy{
				ProcessID:          "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CustomerID:         "19640120-3887",
				PersonalEconomyID:  "52a50f80-3d28-11e9-b210-d663bd873d93",
				YearlyIncome:       340000,
				Income:             0,
				TypeOfEmployeement: "Tillsvidare",
				Employeer:          "Anna Andersson Skog och djurhållning",
				EmployeedFromYear:  "2012"})
		personaleconomies = append(personaleconomies,
			personalEconomy{
				ProcessID:          "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CustomerID:         "19650705-5579",
				PersonalEconomyID:  "bf5ea49c-3d28-11e9-b210-d663bd873d93",
				YearlyIncome:       0,
				Income:             460000,
				TypeOfEmployeement: "Tillsvidare",
				Employeer:          "Katrineholm Revision AB",
				EmployeedFromYear:  "2009"})
	}
	//
	if err := json.NewEncoder(w).Encode(personaleconomies); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}

//
// GetPersonalEconomy
func GetPersonalEconomy(w http.ResponseWriter, r *http.Request) {

	var personaleconomies []personalEconomy
	//
	//vars := mux.Vars(r)
	processid := r.Header.Get("X-process-Id")
	//
	var data personalEconomyID
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
	fmt.Printf("getPersonalEconomies executed: processId: %s /personalEconomyId: %s...\r\n",
		processid, data.PersonalEconomyID)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":

		switch data.PersonalEconomyID {
		case "52a50f80-3d28-11e9-b210-d663bd873d93":
			personaleconomies = append(personaleconomies,
				personalEconomy{
					ProcessID:          "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
					CustomerID:         "19640120-3887",
					PersonalEconomyID:  "52a50f80-3d28-11e9-b210-d663bd873d93",
					YearlyIncome:       340000,
					Income:             0,
					TypeOfEmployeement: "Tillsvidare",
					Employeer:          "Anna Andersson Skog och djurhållning",
					EmployeedFromYear:  "2012"})
		}

		switch data.PersonalEconomyID {
		case "bf5ea49c-3d28-11e9-b210-d663bd873d93":
			personaleconomies = append(personaleconomies,
				personalEconomy{
					ProcessID:          "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
					CustomerID:         "19650705-5579",
					PersonalEconomyID:  "bf5ea49c-3d28-11e9-b210-d663bd873d93",
					YearlyIncome:       0,
					Income:             460000,
					TypeOfEmployeement: "Tillsvidare",
					Employeer:          "Katrineholm Revision AB",
					EmployeedFromYear:  "2009"})
		}
	}
	//
	if err := json.NewEncoder(w).Encode(personaleconomies); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}
