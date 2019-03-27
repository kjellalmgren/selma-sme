package kycinformations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"selmasme/models"

	"github.com/gorilla/mux"
)

// KycEntry documents
func KycEntry(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	switch r.Method {
	case "GET":
		getKycInformations(w, r)
	case "DELETE":
		deleteKycInformation(w, r)
	case "PATCH":
		updateKycInformation(w, r)
	case "PUT":
		addKycInformation(w, r)
	}
}

// GetKycInformations documentation
func getKycInformations(w http.ResponseWriter, r *http.Request) {

	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("getKycInformations executed: processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	file, err := ioutil.ReadFile("kycinformations.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading kycinformtions.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	kycinformations := []models.KycInformation{}
	_ = json.Unmarshal([]byte(file), &kycinformations)
	//
	if err := json.NewEncoder(w).Encode(kycinformations); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

// GetKycInformation information
func getKycInformation(w http.ResponseWriter, r *http.Request) {

	var kycinformations []models.KycInformation
	//
	vars := mux.Vars(r)
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("getKycInformation executed: processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	var data models.CustomerID
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
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		switch data.CustomerID {
		case "19640120-3887":
			switch vars["kycId"] {
			case "498b538c-3d82-11e9-b210-d663bd873d93":
				kycinformations = append(kycinformations,
					models.KycInformation{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
						CustomerID:               "19640120-3887",
						KycID:                    "498b538c-3d82-11e9-b210-d663bd873d93",
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

// deleteKycInformation documentation
func deleteKycInformation(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("deleteKycInformation executed: processId: %s...\r\n", processid)

}

// addKycInformation documentation
func addKycInformation(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("addKycInformation executed: processId: %s...\r\n", processid)
}

// updateKycInformation documentation
func updateKycInformation(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("updateKycInformation executed: processId: %s...\r\n", processid)
}
