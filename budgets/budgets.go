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
	budgets := []models.BudgetType{}
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		file, err := ioutil.ReadFile("json/budgets.json")
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

// GetBudget documentation
func GetBudget(w http.ResponseWriter, r *http.Request) {

	var budgets []models.BudgetType
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	var data models.CompanyEconomyIdBudgetID
	var r1 []byte
	r1, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	//
	//data.BudgetID
	//data.CompanyEconomyID
	json.NewDecoder(bytes.NewReader([]byte(r1))).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("GetBudget executed: processId: %s\r\ncompanyEconomy: %s\r\nBudgetId: %s\r\n",
		processid, data.CompanyEconomyID, data.BudgetID)
	//
	file, err := ioutil.ReadFile("json/budgets.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading budgets.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	budret := make([]models.ValueYearType, 1, 1)
	_ = json.Unmarshal([]byte(file), &budgets)
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		for i := 0; i < len(budgets); i++ {
			if budgets[i].CompanyEconomyID == data.CompanyEconomyID {
				j := 0
				for _, budgetyear := range budgets[i].BudgetYears {
					if budgetyear.BudgetID == data.BudgetID {
						budret[0] = budgets[i].BudgetYears[j]
					}
					j++
				}
			}
		}
	}
	//fmt.Println(fmt.Sprintf("Idx: %s", budret[0].CompanyEconomyID))
	if budret[0].BudgetID == data.BudgetID {
		if err := json.NewEncoder(w).Encode(budret); err != nil {
			w.WriteHeader(http.StatusNotFound)
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Budget not Found", 404)
		//w.WriteHeader(http.StatusNotFound)
	}
}

// UpdateBudget documentation
func UpdateBudget(w http.ResponseWriter, r *http.Request) {

	var budgets []models.ValueYearType

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "PATCH, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	var data models.ValueYearType
	var r1 []byte
	r1, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewDecoder(bytes.NewReader([]byte(r1))).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Println(fmt.Sprintf("BudgetId updated: %v", data.BudgetID))
	for _, budget := range budgets {
		if budget.BudgetID == data.BudgetID {
			fmt.Println(fmt.Sprintf("update year: %d", budget.BudgetYear))
		}
	}
	w.WriteHeader(http.StatusOK)

}

// DeleteBudget documentation
func DeleteBudget(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	var data models.CompanyEconomyIdBudgetID
	var r1 []byte
	r1, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewDecoder(bytes.NewReader([]byte(r1))).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("DeleteBudget executed: processId: %s\r\nCompanyEconomyId: %s \r\nbudgetId: %s\r\n",
		processid, data.CompanyEconomyID, data.BudgetID)
	w.WriteHeader(http.StatusOK)
}

// AddBudget documentation
func AddBudget(w http.ResponseWriter, r *http.Request) {

	var budgets []models.ValueYearType

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	var data models.ValueYearType
	var r1 []byte
	r1, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewDecoder(bytes.NewReader([]byte(r1))).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Println(fmt.Sprintf("BudgetId added: %v", data.BudgetID))
	for _, budget := range budgets {
		if budget.BudgetID == data.BudgetID {
			fmt.Println(fmt.Sprintf("Add year: %d", budget.BudgetYear))
		}
	}
	w.WriteHeader(http.StatusOK)
}
