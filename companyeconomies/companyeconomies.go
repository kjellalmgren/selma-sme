package companyeconomies

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//
// CompanyEconomies
type CompanyEconomy struct {
	ProcessID        string `json:"processId"`
	CompanyID        string `json:"companyId"`
	CompanyEconomyID string `json:"companyEconomyId"`
}

//
// GetCompanyEconomies
func GetCompanyEconomies(w http.ResponseWriter, r *http.Request) {

	var companyeconomies []CompanyEconomy
	//
	vars := mux.Vars(r)
	fmt.Printf("getCompanyEconomies executed: processId: %s...\r\n", vars["processId"])
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//
	switch vars["processId"] {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		companyeconomies = append(companyeconomies,
			CompanyEconomy{
				ProcessID:        "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CompanyID:        "02d6a03e-5895-4077-98f7-7a5c192868b7",
				CompanyEconomyID: "4804f0c2-3d2d-11e9-b210-d663bd873d93"})
	}
	//
	if err := json.NewEncoder(w).Encode(companyeconomies); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}