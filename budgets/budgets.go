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
	fmt.Printf("GetBudgets executed: processId: %s...\r\n", processid)
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		budgets = append(budgets,
			models.Budget{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
				CompanyEconomyID: "4804f0c2-3d2d-11e9-b210-d663bd873d93",
				BudgetID:         "461460c2-3d14-11e9-b210-d663bd873d94",
				BudgetYear:       2018,
				Value1:           200000,
				Value2:           100000,
				Value3:           200000,
				Value4:           200000,
				Value5:           200000,
				Value6:           200000,
				Value7:           200000,
				Value8:           200000,
				Value9:           200000,
				Value10:          200000,
				Value11:          10000,
				Value12:          10000,
				Value13:          300000,
				Value14:          20000,
				Value15:          30000,
				Value16:          200000,
				Value17:          20000,
				Value18:          20000,
				Value19:          20000,
				Value20:          20000,
				Value21:          20000,
				Value22:          20000,
				Value23:          20000,
				Value24:          20000,
				Value25:          20000,
			})
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
		switch data.BudgetID {
		case "461460c2-3d14-11e9-b210-d663bd873d94":
			budgets = append(budgets,
				models.Budget{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
					CompanyEconomyID: "4804f0c2-3d2d-11e9-b210-d663bd873d93",
					BudgetID:         "461460c2-3d14-11e9-b210-d663bd873d94",
					BudgetYear:       2019,
					Value1:           200000,
					Value2:           100000,
					Value3:           200000,
					Value4:           200000,
					Value5:           200000,
					Value6:           200000,
					Value7:           200000,
					Value8:           200000,
					Value9:           200000,
					Value10:          200000,
					Value11:          10000,
					Value12:          10000,
					Value13:          300000,
					Value14:          20000,
					Value15:          30000,
					Value16:          200000,
					Value17:          20000,
					Value18:          20000,
					Value19:          20000,
					Value20:          20000,
					Value21:          20000,
					Value22:          20000,
					Value23:          20000,
					Value24:          20000,
					Value25:          20000,
				})
			budgets = append(budgets,
				models.Budget{ProcessID: "9a65d28a-46bb-4442-b96d-6a09fda6b18b",
					CompanyEconomyID: "4804f0c2-3d2d-11e9-b210-d663bd873d93",
					BudgetID:         "461460c2-3d14-11e9-b210-d663bd873d94",
					BudgetYear:       2020,
					Value1:           200000,
					Value2:           100000,
					Value3:           200000,
					Value4:           200000,
					Value5:           200000,
					Value6:           200000,
					Value7:           200000,
					Value8:           200000,
					Value9:           200000,
					Value10:          200000,
					Value11:          10000,
					Value12:          10000,
					Value13:          300000,
					Value14:          20000,
					Value15:          30000,
					Value16:          200000,
					Value17:          20000,
					Value18:          20000,
					Value19:          20000,
					Value20:          20000,
					Value21:          20000,
					Value22:          20000,
					Value23:          20000,
					Value24:          20000,
					Value25:          20000,
				})
		}
	}
	if err := json.NewEncoder(w).Encode(budgets); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}
