package companies

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Company
type Company struct {
	ProcessID     string `json:"processId,omitempty"`
	CompanyID     string `json:"companyId,omitempty"`
	OrgNumber     string `json:"orgNr,omitempty"`
	CompanyName   string `json:"companyName,omitempty"`
	NumberOfLoans string `json:"numberOfLoans,omitempty"`
	Created       string `json:"created"`
}

// GetCompanies
func GetCompanies(w http.ResponseWriter, r *http.Request) {

	var companies []Company
	//
	vars := mux.Vars(r)
	fmt.Printf("getCompanies executed: processId: %s...\r\n", vars["processId"])
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//
	switch vars["processId"] {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		companies = append(companies,
			Company{
				ProcessID:   "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CompanyID:   "461460c2-3d14-11e9-b210-d663bd873d93",
				OrgNumber:   "551010-8474",
				CompanyName: "Anna Andersson Skog och djurhållning",
				Created:     "2012-01-01"})
	}
	//
	if err := json.NewEncoder(w).Encode(companies); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}

//
// GetCompany
func GetCompany(w http.ResponseWriter, r *http.Request) {

	var companies []Company
	vars := mux.Vars(r)
	fmt.Printf("getCompanies executed: processId: %s...\r\n", vars["processId"])
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//
	switch vars["processId"] {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		switch vars["companyId"] {
		case "461460c2-3d14-11e9-b210-d663bd873d93":
			companies = append(companies,
				Company{
					ProcessID:   "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
					CompanyID:   "461460c2-3d14-11e9-b210-d663bd873d93",
					OrgNumber:   "551010-8474",
					CompanyName: "Anna Andersson Skog och djurhållning",
					Created:     "2012-01-01"})
		}
	}
	//
	if err := json.NewEncoder(w).Encode(companies); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}
