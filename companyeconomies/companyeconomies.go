package companyeconomies

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//
// CompanyEconomies
type companyEconomy struct {
	ProcessID        string `json:"processId"`
	CompanyID        string `json:"companyId"`
	CompanyEconomyID string `json:"companyEconomyId"`
}

//
// GetCompanyEconomies
func GetCompanyEconomies(w http.ResponseWriter, r *http.Request) {

	var companyeconomies []companyEconomy
	//
	vars := mux.Vars(r)
	fmt.Printf("getCompanyEconomies executed: processId: %s...\r\n", vars["processId"])
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	switch vars["processId"] {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		companyeconomies = append(companyeconomies,
			companyEconomy{
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
func CompanyEconomyEntry(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	switch r.Method {
	case "POST":
		getCompanyEconomy(w, r)
	case "DELETE":
		deleteCompanyEconomy(w, r)
	case "PATCH":
		updateCompanyEconomy(w, r)
	}
}

// GetCompanyEconomy
func getCompanyEconomy(w http.ResponseWriter, r *http.Request) {

	var companyeconomies []companyEconomy

	//
	vars := mux.Vars(r)
	fmt.Printf("getCompanyEconomy executed: processId: %s / companyId: %s / companyEconomyId: %s...\r\n",
		vars["processId"], vars["companyId"], vars["companyEconomyId"])
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	switch vars["processId"] {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		switch vars["companyEconomyId"] {
		case "4804f0c2-3d2d-11e9-b210-d663bd873d93":
			companyeconomies = append(companyeconomies,
				companyEconomy{
					ProcessID:        "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
					CompanyID:        "02d6a03e-5895-4077-98f7-7a5c192868b7",
					CompanyEconomyID: "4804f0c2-3d2d-11e9-b210-d663bd873d93"})
		}
	}
	//
	if err := json.NewEncoder(w).Encode(companyeconomies); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}

func deleteCompanyEconomy(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	//varsh := r.Header
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("deleteCompanyEconomy executed, processId: %s...\r\n", processid)
	w.WriteHeader(http.StatusOK)
}
func updateCompanyEconomy(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	//varsh := r.Header
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("updateCompanyEconomy executed, processId: %s...\r\n", processid)
	w.WriteHeader(http.StatusOK)
}
