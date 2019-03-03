package collaterals

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Collaterals
type Collateral struct {
	ProcessID              string `json:"processId,omitempty"`
	CustomerID             string `json:"customerId,omitempty"`
	CollateralID           string `json:"collateralId,omitempty"`
	CollateralName         string `json:"collateralName,omitempty"`
	CollateralMunicipality string `json:"collateralMunicipality,omitempty"`
	CollateralStreet       string `json:"collateralStreet,omitempty"`
	UseAsCollateral        bool   `json:"useAsCollateral"`
	BuyCollateral          bool   `json:"buyCollateral"`
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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//
	switch vars["processId"] {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		collaterals = append(collaterals,
			Collateral{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CustomerID:             "19640120-3887",
				CollateralID:           "b25961de-3cc3-11e9-b210-d663bd873d93",
				CollateralName:         "ÅGERSTA 1:6",
				CollateralMunicipality: "ENKÖPING",
				CollateralStreet:       "Bergsgatan 12",
				UseAsCollateral:        false,
				BuyCollateral:          true})

	}
	if err := json.NewEncoder(w).Encode(collaterals); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}
