package collaterals

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Collaterals
type Collateral struct {
	ProcessID              string `json:"processId"`
	CustomerID             string `json:"customerId"`
	CollateralID           string `json:"collateralId"`
	CollateralCode         string `json:"collateralCode"`
	CollateralName         string `json:"collateralName"`
	TaxedOwners            []TaxedOwnerType
	CollateralMunicipality string `json:"collateralMunicipality"`
	CollateralStreet       string `json:"collateralStreet"`
	UseAsCollateral        bool   `json:"useAsCollateral"`
	BuyCollateral          bool   `json:"buyCollateral"`
}

type TaxedOwnerType struct {
	TaxedId    string `json:"taxedId"`
	TaxedOwner string `json:"taxedOwner"`
}

// GetCollaterals
func GetCollaterals(w http.ResponseWriter, r *http.Request) {

	var collaterals []Collateral
	//
	vars := mux.Vars(r)
	fmt.Printf("getCollaterals executed: processId: %s...\r\n", vars["processId"])
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
			Collateral{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CustomerID:     "19640120-3887",
				CollateralID:   "b25961de-3cc3-11e9-b210-d663bd873d93",
				CollateralCode: "10",
				CollateralName: "ÅGERSTA 1:6",
				TaxedOwners: []TaxedOwnerType{
					TaxedOwnerType{
						TaxedId:    "c73119bc-3e71-11e9-b210-d663bd873d93",
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

	var collaterals []Collateral
	//
	vars := mux.Vars(r)
	fmt.Printf("getCollateral executed: processId: %s /collateralId: %s...\r\n",
		vars["processId"], vars["collateralId"])
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	switch vars["processId"] {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		switch vars["collateralId"] {
		case "b25961de-3cc3-11e9-b210-d663bd873d93":
			collaterals = append(collaterals,
				Collateral{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
					CustomerID:     "19640120-3887",
					CollateralID:   "b25961de-3cc3-11e9-b210-d663bd873d93",
					CollateralCode: "10",
					CollateralName: "ÅGERSTA 1:6",
					TaxedOwners: []TaxedOwnerType{
						TaxedOwnerType{
							TaxedId:    "c73119bc-3e71-11e9-b210-d663bd873d93",
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
