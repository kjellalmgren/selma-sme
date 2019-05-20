package processes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"selmasme/models"
)

//
// ProcessEntry documentation
func processEntry(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID, caseIdStatus")
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//
	switch r.Method {
	case "DELETE":
		DeleteProcess(w, r)
	case "PUT":
		AddProcess(w, r)
	}
}

//
// ***************************************************************
//
// GetProcesses documentation
func GetProcesses(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID, X-customer-ID")
	//
	//processid := r.Header.Get("X-process-Id")
	//fmt.Printf("getProcesses executed, processid: %s...\r\n", processid)
	//
	//var data customerID
	//customerid := r.Header.Get("X-customer-Id")
	//fmt.Printf("getProcesses executed, customerID: %s...\r\n", customerid)
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
	fmt.Printf("getProcesses executed, customerID: %s...\r\n", data.CustomerID)
	//
	file, err := ioutil.ReadFile("json/processes.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading processes.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	processes := []models.ProcessType{}
	_ = json.Unmarshal([]byte(file), &processes)
	//
	//
	//for i, p := range processes {
	//	fmt.Println(i, p.ProcessID, p.CustomerID)
	//}
	//
	if err := json.NewEncoder(w).Encode(processes); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

//
// DeleteProcess documentation
func DeleteProcess(w http.ResponseWriter, r *http.Request) {

	processid := r.Header.Get("X-process-Id")
	fmt.Printf("deleteProcess executed, processid: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID, caseIdStatus")
	//
	caseidstatus := r.Header.Get("caseIdStatus")
	fmt.Printf("deleteProcess executed, caseIdStatus: %s...\r\n", caseidstatus)
	w.WriteHeader(http.StatusOK)
	//
}

// AddProcess documentation
func AddProcess(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID, caseIdStatus")
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
	fmt.Printf("AddProcess executed, will return a processid and a caseId from Loan Process for customer: %s...\r\n", data.CustomerID)

	// *** Status has to be setr i production ***
	//caseidstatus := r.Header.Get("caseIdStatus")
	processidcaseid := models.ProcessIDCaseID{}
	//
	processidcaseid.ProcessID = "9a65d28a-46bb-4442-b96d-6a09fda6b18b"
	processidcaseid.CaseID = "106100"
	//
	if err := json.NewEncoder(w).Encode(processidcaseid); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
}

// UpdateProcess documentation
func UpdateProcess(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "PATCH, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID, caseIdStatus")
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
	fmt.Printf("UpdateProcess executed, will return a processid and a caseId from Loan Process for customer: %s...\r\n", data.CustomerID)

	// *** Status has to be setr i production ***
	//caseidstatus := r.Header.Get("caseIdStatus")
	processidcaseid := models.ProcessIDCaseID{}
	//
	processidcaseid.ProcessID = "9a65d28a-46bb-4442-b96d-6a09fda6b18b"
	processidcaseid.CaseID = "106100"
	//
	if err := json.NewEncoder(w).Encode(processidcaseid); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

// SaveProcessAll documentation
func SaveProcessAll(w http.ResponseWriter, r *http.Request) {

	processid := r.Header.Get("X-process-Id")
	fmt.Printf("SaveProcessAll executed, processid: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID, caseIdStatus")
	//
	var data models.ProcessAllType
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
	fmt.Println(fmt.Sprintf("# of Applicant:%v", len(data.Applicants)))
	//
	w.WriteHeader(http.StatusOK)
	//
}

// GetProcessAll documentation
func GetProcessAll(w http.ResponseWriter, r *http.Request) {
	//
	processall := models.ProcessAllType{}
	processes := []models.ProcessType{}
	applicants := []models.ApplicantType{}
	loans := []models.LoanType{}
	extloans := []models.ExtLoanType{}
	households := []models.HouseholdType{}
	companies := []models.CompanyType{}
	companyeconomies := []models.CompanyEconomyType{}
	personaleconomies := []models.PersonalEconomyType{}
	kycinformations := []models.KycInformationType{}
	collaterals := []models.CollateralType{}
	budgets := []models.BudgetType{}
	eusupports := []models.EUSupportType{}
	guarantors := []models.GuarantorType{}
	maintenancecosts := []models.MaintenanceCostType{}
	//
	//var d models.XAll
	//var p models.Person
	//p.Name = "Kjell Almgren"
	//p.Mobile = "0733981482"
	//d.Persons = append(d.Persons, p)
	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("GetProcessAll executed, processid: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID, caseIdStatus")
	//
	file, err := ioutil.ReadFile("json/processes.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading processes.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	_ = json.Unmarshal([]byte(file), &processes)
	//
	file, err = ioutil.ReadFile("json/applicants.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading applicants.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	_ = json.Unmarshal([]byte(file), &applicants)
	//
	file, err = ioutil.ReadFile("json/loans.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading loans.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	_ = json.Unmarshal([]byte(file), &loans)
	file, err = ioutil.ReadFile("json/companies.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading companies.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	//companies = []models.Company{}
	_ = json.Unmarshal([]byte(file), &companies)
	//
	file, err = ioutil.ReadFile("json/personaleconomies.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading personaleconomies.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	_ = json.Unmarshal([]byte(file), &personaleconomies)
	//
	file, err = ioutil.ReadFile("json/companyeconomies.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading companyeconomies.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	_ = json.Unmarshal([]byte(file), &companyeconomies)
	//
	file, err = ioutil.ReadFile("json/extloans.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading extloans.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	_ = json.Unmarshal([]byte(file), &extloans)
	//
	file, err = ioutil.ReadFile("json/households.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading households.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	//households = []models.Household{}
	_ = json.Unmarshal([]byte(file), &households)
	//
	file, err = ioutil.ReadFile("json/collaterals.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading collaterals.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	_ = json.Unmarshal([]byte(file), &collaterals)
	//
	file, err = ioutil.ReadFile("json/budgets.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading budgets.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	_ = json.Unmarshal([]byte(file), &budgets)
	//
	file, err = ioutil.ReadFile("json/kycinformations.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading kycinformations.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	_ = json.Unmarshal([]byte(file), &kycinformations)
	//
	file, err = ioutil.ReadFile("json/eusupports.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading eusupports.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	_ = json.Unmarshal([]byte(file), &eusupports)
	//
	file, err = ioutil.ReadFile("json/guarantors.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading guarantors.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	_ = json.Unmarshal([]byte(file), &guarantors)
	//
	file, err = ioutil.ReadFile("json/mainteancecosts.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading mainteancecosts.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	_ = json.Unmarshal([]byte(file), &maintenancecosts)
	//
	processall.Processes = append(processes)
	processall.Applicants = append(applicants)
	processall.Loans = append(loans)
	processall.ExtLoans = append(extloans)
	processall.Households = append(households)
	processall.Companies = append(companies)
	processall.CompanyEconomies = append(companyeconomies)
	processall.PersonalEconomies = append(personaleconomies)
	processall.KycInformations = append(kycinformations)
	processall.Collaterals = append(collaterals)
	processall.Budgets = append(budgets)
	processall.EUSupports = append(eusupports)
	processall.Guarantors = append(guarantors)
	processall.MaintenanceCosts = append(maintenancecosts)
	//
	if err := json.NewEncoder(w).Encode(processall); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}
