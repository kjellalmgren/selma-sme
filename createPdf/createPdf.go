package createpdf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"selmasme/models"
	"time"

	"github.com/jung-kurt/gofpdf"
)

// TableRow documentation
type TableRow struct {
	Key   string
	Value string
}

// HeaderRow
type HeaderRow struct {
	Value string
}

// CreatePdfDocument documentation
func CreatePdfDocument(processid string) models.MessageBody {

	message := models.MessageBody{}
	//
	processall := models.ProcessAllType{}
	processes := getProcesses(processid)
	applicants := getApplicants(processid)
	companies := getCompanies(processid)
	collaterals := getCollaterals(processid)
	loans := getLoans(processid)
	extloans := getExtLoans(processid)
	households := getHouseholds(processid)
	companyeconomies := getCompanyEconomies(processid)
	personaleconomies := getPersonalEconomies(processid)
	kycinformations := getKycInformations(processid)
	budgets := getBudgets(processid)
	//
	processall.Processes = append(processes)
	processall.Applicans = append(applicants)
	processall.Companies = append(companies)
	processall.Collaterals = append(collaterals)
	processall.Loans = append(loans)
	processall.ExtLoans = append(extloans)
	processall.Households = append(households)
	processall.CompanyEconomies = append(companyeconomies)
	processall.PersonalEconomies = append(personaleconomies)
	processall.KycInformations = append(kycinformations)
	processall.Budgets = append(budgets)
	//
	// ##########################################
	// # Create the actual pdf document
	// ##########################################
	t := time.Now().Local()
	fmt.Println(t.String())
	fmt.Println(t.Format("2006-01-02 15:04:05"))

	filename := processid + "-" + t.Format("2006-01-02 15:04:05") + ".pdf"
	fmt.Println(filename)
	//
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	tr := []TableRow{} // Initiate table on page one
	hr := HeaderRow{}
	// Processes
	pdf.SetFont("Arial", "B", 16)
	hr.Value = "Processes"
	//
	for _, process := range processes {
		pdf = header1(pdf, hr.Value)
		pdf.SetFont("Arial", "B", 11)
		tr = append(tr, TableRow{Key: "ProcessID:", Value: process.ProcessID})
		tr = append(tr, TableRow{Key: "Created date:", Value: process.ProcessCreatedDate})
		for _, customerID := range process.CustomerID {
			tr = append(tr, TableRow{Key: "CustomerId:", Value: customerID.CustomerID})
		}
	}
	pdf = table1(pdf, tr) // add table to page
	// Applicants
	tr = []TableRow{}
	hr = HeaderRow{}
	pdf.SetFont("Arial", "B", 16)
	hr.Value = "Applicants"
	pdf.SetFont("Arial", "B", 11)
	for _, applicant := range applicants {
		pdf = header1(pdf, hr.Value)
		tr = append(tr, TableRow{Key: "customerId:", Value: applicant.CustomerID})
		tr = append(tr, TableRow{Key: "name:", Value: applicant.ApplicantName})
		tr = append(tr, TableRow{Key: "Adress:", Value: applicant.ApplicantAddress})
		tr = append(tr, TableRow{Key: "PostNr:", Value: applicant.ApplicantPostAddress})
		for _, ci := range applicant.ContactInformation {
			log.Printf("eMail: %s - MobileNumber: %s", ci.ApplicantEmail, ci.ApplicantMobileNumber)
			tr = append(tr, TableRow{Key: "eMail:", Value: ci.ApplicantEmail})
			tr = append(tr, TableRow{Key: "phone:", Value: ci.ApplicantMobileNumber})
		}
		tr = append(tr, TableRow{Key: "Stakeholder:", Value: applicant.StakeholderType})
		tr = append(tr, TableRow{Key: "Kontakt via sms:", Value: fmt.Sprintf("%v", applicant.ApplicantBySms)})
		tr = append(tr, TableRow{Key: "Kontakt via eMail:", Value: fmt.Sprintf("%v", applicant.ApplicantByeMail)})
		tr = append(tr, TableRow{Key: "Anställd:", Value: fmt.Sprintf("%v", applicant.ApplicantEmployeed)})
	}
	pdf = table1(pdf, tr) // add table to page current page

	// Companies
	//
	tr = []TableRow{}
	hr = HeaderRow{}
	pdf.SetFont("Arial", "B", 16)
	hr.Value = "Companies"
	pdf.SetFont("Arial", "B", 11)
	for _, company := range companies {
		pdf = header1(pdf, hr.Value)
		tr = append(tr, TableRow{Key: "companyId:", Value: company.CompanyID})
		tr = append(tr, TableRow{Key: "OrgNr:", Value: company.OrgNumber})
		tr = append(tr, TableRow{Key: "CompanyName:", Value: company.CompanyName})
		tr = append(tr, TableRow{Key: "BusinessFocus:", Value: company.BusinessFocus})
		tr = append(tr, TableRow{Key: "Created:", Value: company.Created})
		tr = append(tr, TableRow{Key: "SNI-code:", Value: company.IndustriCode})
		tr = append(tr, TableRow{Key: "SNI-Text:", Value: company.IndustriText})
		tr = append(tr, TableRow{Key: "Selected:", Value: fmt.Sprintf("%v", company.SelectedCompany)})

	}
	pdf = table1(pdf, tr) // add table to page current page
	//
	// Collaterals
	//
	tr = []TableRow{}
	hr = HeaderRow{}
	pdf.SetFont("Arial", "B", 16)
	hr.Value = "Collaterals"
	for _, collateral := range collaterals {
		pdf = header1(pdf, hr.Value)
		pdf.SetFont("Arial", "B", 11)
		tr = append(tr, TableRow{Key: "collateralID:", Value: collateral.CollateralID})
		tr = append(tr, TableRow{Key: "Fastighetskod:", Value: collateral.CollateralCode})
		tr = append(tr, TableRow{Key: "CustomerID:", Value: collateral.CustomerID})
		tr = append(tr, TableRow{Key: "ProcessID:", Value: collateral.ProcessID})
		tr = append(tr, TableRow{Key: "Kommun:", Value: collateral.CollateralMunicipality})
		tr = append(tr, TableRow{Key: "fastighetsnamn", Value: collateral.CollateralName})
		tr = append(tr, TableRow{Key: "Gata:", Value: collateral.CollateralStreet})
		tr = append(tr, TableRow{Key: "Använd säkerhet:", Value: fmt.Sprintf("%v", collateral.UseAsCollateral)})
		tr = append(tr, TableRow{Key: "Köpa säkerhet:", Value: fmt.Sprintf("%v", collateral.BuyCollateral)})
		for _, taxedOwner := range collateral.TaxedOwners {
			tr = append(tr, TableRow{Key: "TaxedID:", Value: taxedOwner.TaxedID})
			tr = append(tr, TableRow{Key: "Owner:", Value: taxedOwner.TaxedOwner})
		}
	}
	//
	// Budget
	//
	tr = []TableRow{}
	hr = HeaderRow{}
	pdf.SetFont("Arial", "B", 16)
	hr.Value = "Budget"
	pdf.SetFont("Arial", "B", 11)
	for _, budget := range budgets {
		pdf = header1(pdf, hr.Value)
		tr = append(tr, TableRow{Key: "CompanyID:", Value: budget.CompanyEconomyID})
		for _, budgetyear := range budget.BudgetYears {
			tr = append(tr, TableRow{Key: "Year:", Value: fmt.Sprintf("%v", budgetyear.BudgetYear)})
			tr = append(tr, TableRow{Key: "Value1:", Value: fmt.Sprintf("%v", budgetyear.Value1)})
			tr = append(tr, TableRow{Key: "Value2:", Value: fmt.Sprintf("%v", budgetyear.Value2)})
			tr = append(tr, TableRow{Key: "Value3:", Value: fmt.Sprintf("%v", budgetyear.Value3)})
			tr = append(tr, TableRow{Key: "Value4:", Value: fmt.Sprintf("%v", budgetyear.Value4)})
			tr = append(tr, TableRow{Key: "Value5:", Value: fmt.Sprintf("%v", budgetyear.Value5)})
			tr = append(tr, TableRow{Key: "Value6:", Value: fmt.Sprintf("%v", budgetyear.Value6)})
			tr = append(tr, TableRow{Key: "Value7:", Value: fmt.Sprintf("%v", budgetyear.Value7)})
			tr = append(tr, TableRow{Key: "Value8:", Value: fmt.Sprintf("%v", budgetyear.Value8)})
			tr = append(tr, TableRow{Key: "Value9:", Value: fmt.Sprintf("%v", budgetyear.Value9)})
			tr = append(tr, TableRow{Key: "Value10:", Value: fmt.Sprintf("%v", budgetyear.Value10)})
			tr = append(tr, TableRow{Key: "Value11:", Value: fmt.Sprintf("%v", budgetyear.Value11)})
			tr = append(tr, TableRow{Key: "Value12:", Value: fmt.Sprintf("%v", budgetyear.Value12)})
			tr = append(tr, TableRow{Key: "Value13:", Value: fmt.Sprintf("%v", budgetyear.Value13)})
			tr = append(tr, TableRow{Key: "Value14:", Value: fmt.Sprintf("%v", budgetyear.Value14)})
			tr = append(tr, TableRow{Key: "Value15:", Value: fmt.Sprintf("%v", budgetyear.Value15)})
			tr = append(tr, TableRow{Key: "Value16:", Value: fmt.Sprintf("%v", budgetyear.Value16)})
			tr = append(tr, TableRow{Key: "Value17:", Value: fmt.Sprintf("%v", budgetyear.Value17)})
			tr = append(tr, TableRow{Key: "Value18:", Value: fmt.Sprintf("%v", budgetyear.Value18)})
			tr = append(tr, TableRow{Key: "Value19:", Value: fmt.Sprintf("%v", budgetyear.Value19)})
			tr = append(tr, TableRow{Key: "Value20:", Value: fmt.Sprintf("%v", budgetyear.Value20)})
			tr = append(tr, TableRow{Key: "Value21:", Value: fmt.Sprintf("%v", budgetyear.Value21)})
			tr = append(tr, TableRow{Key: "Value22:", Value: fmt.Sprintf("%v", budgetyear.Value22)})
			tr = append(tr, TableRow{Key: "Value23:", Value: fmt.Sprintf("%v", budgetyear.Value23)})
			tr = append(tr, TableRow{Key: "Value24:", Value: fmt.Sprintf("%v", budgetyear.Value24)})
			tr = append(tr, TableRow{Key: "Value25:", Value: fmt.Sprintf("%v", budgetyear.Value25)})

		}
	}
	//
	pdf = table1(pdf, tr) // add table to page current page
	//
	// Save pdf file to a local destination
	//
	err := savePDF(filename, pdf)
	if err != nil {
		log.Fatalf("createPdf::CreatePdfDocument: Cannot save PDF: %s|n", filename)
		message.Status = false
		message.MessageText = "Error creating pdf document..."
		message.Filename = ""
	} else {
		message.Status = true
		message.MessageText = "OK"
		message.Filename = ""
	}
	return message
}

//
// getProcesses documentation
func getProcesses(processid string) []models.ProcessType {

	processes := []models.ProcessType{}
	//procret := make([]models.ProcessType, 1, 3)
	var procret []models.ProcessType

	file, err := ioutil.ReadFile("json/processes.json")
	if err != nil {
		fmt.Println("Error reading processes.json - %s", err)
	}
	_ = json.Unmarshal([]byte(file), &processes)
	//j := 0
	for _, process := range processes {

		if process.ProcessID == processid {
			//procret[j] = process
			procret = append(procret, process)
			//j++
		}
	}
	return procret
}

//
// getApplicants documentation
func getApplicants(processid string) []models.ApplicantType {

	//message := models.MessageBody{}
	applicants := []models.ApplicantType{}
	//appret := make([]models.ApplicantType, 2, 2)
	var appret []models.ApplicantType
	//
	file, err := ioutil.ReadFile("json/applicants.json")
	if err != nil {
		fmt.Println("Error reading applicants.json - %s", err)
	}
	_ = json.Unmarshal([]byte(file), &applicants)
	//
	//j := 0
	for _, applicant := range applicants {
		if applicant.ProcessID == processid {
			//appret[j] = applicant
			appret = append(appret, applicant)
			//j++
		}
	}
	return appret
}

// getCompanies documentation
func getCompanies(processid string) []models.CompanyType {

	//message := models.MessageBody{}
	companies := []models.CompanyType{}
	var compret []models.CompanyType
	//compret := make([]models.CompanyType, 1, length+1)
	//fmt.Println("Length=%d", length)
	//
	file, err := ioutil.ReadFile("json/companies.json")
	if err != nil {
		fmt.Println("Error reading companies.json - %s", err)
	}
	_ = json.Unmarshal([]byte(file), &companies)
	//
	//j := 0
	for _, company := range companies {
		if company.ProcessID == processid {
			//compret[j] = company
			compret = append(compret, company)
			//j++
		}
	}
	return compret
}

// getCollaterals documentation
func getCollaterals(processid string) []models.CollateralType {

	//message := models.MessageBody{}
	collaterals := []models.CollateralType{}
	var collret []models.CollateralType
	//
	file, err := ioutil.ReadFile("json/collaterals.json")
	if err != nil {
		fmt.Println("Error reading collaterals.json - %s", err)
	}
	_ = json.Unmarshal([]byte(file), &collaterals)
	//
	for _, collateral := range collaterals {
		if collateral.ProcessID == processid {
			//collret[j] = collateral
			collret = append(collret, collateral)
		}
	}
	return collret
}

// getLoans documentation
func getLoans(processid string) []models.LoanType {

	//message := models.MessageBody{}
	loans := []models.LoanType{}
	var loanret []models.LoanType
	//loanret := make([]models.LoanType, 2, 2)
	//
	file, err := ioutil.ReadFile("json/loans.json")
	if err != nil {
		fmt.Println("Error reading loans.json - %s", err)
	}
	_ = json.Unmarshal([]byte(file), &loans)
	//
	for _, loan := range loans {
		if loan.ProcessID == processid {
			//loanret[j] = loan
			loanret = append(loanret, loan)
		}
	}
	return loanret
}

// getExtLoans documentation
func getExtLoans(processid string) []models.ExtLoanType {

	//message := models.MessageBody{}
	extloans := []models.ExtLoanType{}
	//loanret := make([]models.ExtLoanType, 2, 2)
	var loanret []models.ExtLoanType
	//
	file, err := ioutil.ReadFile("json/extloans.json")
	if err != nil {
		fmt.Println("Error reading extloans.json - %s", err)
	}
	_ = json.Unmarshal([]byte(file), &extloans)
	//
	for _, extloan := range extloans {
		if extloan.ProcessID == processid {
			//loanret[j] = extloan
			loanret = append(loanret, extloan)
		}
	}
	return loanret
}

// getHouseholds documentation
func getHouseholds(processid string) []models.HouseholdType {

	//message := models.MessageBody{}
	households := []models.HouseholdType{}
	// householdret := make([]models.HouseholdType, 2, 2)
	var householdret []models.HouseholdType
	//
	file, err := ioutil.ReadFile("json/households.json")
	if err != nil {
		fmt.Println("Error reading households.json - %s", err)
	}
	_ = json.Unmarshal([]byte(file), &households)
	//
	for _, household := range households {
		if household.ProcessID == processid {
			//householdret[j] = household
			householdret = append(householdret, household)
		}
	}
	return householdret
}

// getCompanyEconomies documentation
func getCompanyEconomies(processid string) []models.CompanyEconomyType {

	//message := models.MessageBody{}
	companyeconomies := []models.CompanyEconomyType{}
	//compret := make([]models.CompanyEconomyType, 2, 2)
	var compret []models.CompanyEconomyType
	//
	file, err := ioutil.ReadFile("json/companyeconomies.json")
	if err != nil {
		fmt.Println("Error reading companyeconomies.json - %s", err)
	}
	_ = json.Unmarshal([]byte(file), &companyeconomies)
	//
	for _, companyeconomy := range companyeconomies {
		if companyeconomy.ProcessID == processid {
			//compret[j] = companyeconomy
			compret = append(compret, companyeconomy)
		}
	}
	return compret
}

// getPersonalEconomies documentation
func getPersonalEconomies(processid string) []models.PersonalEconomyType {

	//message := models.MessageBody{}
	personaleconomies := []models.PersonalEconomyType{}
	//persret := make([]models.PersonalEconomyType, 2, 2)
	var persret []models.PersonalEconomyType
	//
	file, err := ioutil.ReadFile("json/personaleconomies.json")
	if err != nil {
		fmt.Println("Error reading personaleconomies.json - %s", err)
	}
	_ = json.Unmarshal([]byte(file), &personaleconomies)
	//
	for _, personaleconomy := range personaleconomies {
		if personaleconomy.ProcessID == processid {
			//persret[j] = personaleconomy
			persret = append(persret, personaleconomy)
		}
	}
	return persret
}

// getKycInformartions documentation
func getKycInformations(processid string) []models.KycInformationType {

	//message := models.MessageBody{}
	kycinformations := []models.KycInformationType{}
	//kycret := make([]models.KycInformationType, 2, 2)
	var kycret []models.KycInformationType
	//
	file, err := ioutil.ReadFile("json/kycinformations.json")
	if err != nil {
		fmt.Println("Error reading kycinformations.json - %s", err)
	}
	_ = json.Unmarshal([]byte(file), &kycinformations)
	//
	for _, kycinformation := range kycinformations {
		if kycinformation.ProcessID == processid {
			//kycret[j] = kycinformation
			kycret = append(kycret, kycinformation)
		}
	}
	return kycret
}

// GetBudgets documentation
func getBudgets(processid string) []models.BudgetType {

	//message := models.MessageBody{}
	budgets := []models.BudgetType{}
	//budret := make([]models.BudgetType, 2, 2)
	var budret []models.BudgetType
	//
	file, err := ioutil.ReadFile("json/budgets.json")
	if err != nil {
		fmt.Println("Error reading budgets.json - %s", err)
	}
	_ = json.Unmarshal([]byte(file), &budgets)
	//
	for _, budget := range budgets {
		if budget.ProcessID == processid {
			//budret[j] = budget
			budret = append(budret, budget)
		}
	}
	return budret
}

// savePDF documentation
func savePDF(filename string, pdf *gofpdf.Fpdf) error {
	return pdf.OutputFileAndClose(filename)
}

// header documentation
func header1(pdf *gofpdf.Fpdf, hdr string) *gofpdf.Fpdf {
	pdf.SetFont("Arial", "B", 16)
	pdf.SetFillColor(240, 240, 240)
	pdf.CellFormat(40, 7, hdr, "1", 0, "", true, 0, "")
	pdf.Ln(-1)
	return pdf
}

// table1 documentation
func table1(pdf *gofpdf.Fpdf, tbl []TableRow) *gofpdf.Fpdf {
	// Reset font and fill color.
	pdf.SetFont("Arial", "", 16)
	pdf.SetFillColor(255, 255, 255)
	// Every column gets aligned according to its contents.
	//align := []string{"L", "C", "L", "R", "R", "R"}
	for _, line := range tbl {
		pdf.CellFormat(50, 10, line.Key, "1", 0, "L", false, 0, "")
		pdf.CellFormat(120, 10, line.Value, "1", 0, "L", false, 0, "")

		pdf.Ln(-1)
	}
	return pdf
}

// image documentation
func image(pdf *gofpdf.Fpdf) *gofpdf.Fpdf {
	//The ImageOptions method takes a file path, x, y, width, and height parameters, and an ImageOptions struct to specify a couple of options.
	pdf.ImageOptions("stats.png", 225, 10, 25, 25, false, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "")
	return pdf
}
