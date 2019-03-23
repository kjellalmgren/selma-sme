package budgets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"selmasme/models"
)

// GetBudgets
func GetBudgets(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("GetBudgets executed: processId: %s...\r\n", processid)
	//
	budgets := []models.Budget{}
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		file, err := ioutil.ReadFile("budgets.json")
		if err != nil {
			fmt.Fprintf(w, "Error reading budgets.json - %s", err)
			w.WriteHeader(http.StatusNotFound)
		}
		_ = json.Unmarshal([]byte(file), &budgets)
	}
	if err := json.NewEncoder(w).Encode(budgets); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

// GetBudget
func GetBudget(w http.ResponseWriter, r *http.Request) {

	var budgets []models.Budget
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("GetBudget executed: processId: %s...\r\n", processid)
	//
	var data models.BudgetID
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

	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":

	}
	if err := json.NewEncoder(w).Encode(budgets); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}
