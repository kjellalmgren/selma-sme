package collaterals

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"selmasme/models"

	"github.com/gorilla/mux"
)

// GetCollaterals
func GetCollaterals(w http.ResponseWriter, r *http.Request) {

	var collaterals []models.Collateral
	//
	vars := mux.Vars(r)
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("getCollaterals executed: processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	switch vars["processId"] {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		collaterals = append(collaterals,
			models.Collateral{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CustomerID:     "19640120-3887",
				CollateralID:   "b25961de-3cc3-11e9-b210-d663bd873d93",
				CollateralCode: "10",
				CollateralName: "ÅGERSTA 1:6",
				TaxedOwners: []models.TaxedOwnerType{
					models.TaxedOwnerType{
						TaxedID:    "c73119bc-3e71-11e9-b210-d663bd873d93",
						TaxedOwner: "Anna Andersson",
					},
				},
				CollateralMunicipality: "ENKÖPING",
				CollateralStreet:       "Bergsgatan 12",
				UseAsCollateral:        false,
				BuyCollateral:          true})
	}
	//
	if err := json.NewEncoder(w).Encode(collaterals); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

// GetCollateral
func GetCollateral(w http.ResponseWriter, r *http.Request) {

	var collaterals []models.Collateral
	//
	processid := r.Header.Get("X-process-Id")
	var data models.CollateralID
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
	fmt.Printf("getCollateral executed: processId: %s /collateralId: %s...\r\n",
		processid, data.CollateralID)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		switch data.CollateralID {
		case "b25961de-3cc3-11e9-b210-d663bd873d93":
			collaterals = append(collaterals,
				models.Collateral{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
					CustomerID:     "19640120-3887",
					CollateralID:   "b25961de-3cc3-11e9-b210-d663bd873d93",
					CollateralCode: "10",
					CollateralName: "ÅGERSTA 1:6",
					TaxedOwners: []models.TaxedOwnerType{
						models.TaxedOwnerType{
							TaxedID:    "c73119bc-3e71-11e9-b210-d663bd873d93",
							TaxedOwner: "Anna Andersson",
						},
					},
					CollateralMunicipality: "ENKÖPING",
					CollateralStreet:       "Bergsgatan 12",
					UseAsCollateral:        false,
					BuyCollateral:          true})
		}
	}
	//
	if err := json.NewEncoder(w).Encode(collaterals); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}
