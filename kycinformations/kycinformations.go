package kycinformations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// KycInformations
type KycInformation struct {
	ProcessID                string `json:"processId"`
	CustomerID               string `json:"customerId"`
	KycId                    string `json:"kycId"`
	KycAcceptUC              bool   `json:"kycAcceptUC"`
	KycAcceptGDPR            bool   `json:"kycAcceptGDPR"`
	KycUCAware               bool   `json:"kycUCAware"`
	KycPublicFunction        bool   `json:"kycPublicFunction"`
	KycRelatedPublicFunction bool   `json:"kycRelatedPublicFunction"`
}

// GetKycInformations
func GetKycInformations(w http.ResponseWriter, r *http.Request) {

	var kycinformations []KycInformation
	//
	vars := mux.Vars(r)
	fmt.Printf("getKycInformations executed: processId: %s...\r\n", vars["processId"])
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	switch vars["processId"] {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		kycinformations = append(kycinformations,
			KycInformation{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CustomerID:               "19640120-3887",
				KycId:                    "498b538c-3d82-11e9-b210-d663bd873d93",
				KycAcceptUC:              true,
				KycAcceptGDPR:            true,
				KycUCAware:               true,
				KycPublicFunction:        false,
				KycRelatedPublicFunction: false})
	}
	//
	if err := json.NewEncoder(w).Encode(kycinformations); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

// GetKycInformation
func GetKycInformation(w http.ResponseWriter, r *http.Request) {

	var kycinformations []KycInformation
	//
	vars := mux.Vars(r)
	fmt.Printf("getKycInformations executed: processId: %s...\r\n", vars["processId"])
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//
	switch vars["processId"] {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		switch vars["customerId"] {
		case "19640120-3887":
			switch vars["kycId"] {
			case "498b538c-3d82-11e9-b210-d663bd873d93":
				kycinformations = append(kycinformations,
					KycInformation{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
						CustomerID:               "19640120-3887",
						KycId:                    "498b538c-3d82-11e9-b210-d663bd873d93",
						KycAcceptUC:              true,
						KycAcceptGDPR:            true,
						KycUCAware:               true,
						KycPublicFunction:        false,
						KycRelatedPublicFunction: false})
			}
		}
	}
	//
	if err := json.NewEncoder(w).Encode(kycinformations); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}
