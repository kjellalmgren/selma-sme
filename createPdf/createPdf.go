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

type BudgetTableRow struct {
	Key     string
	Text    string
	ValueC1 string
	ValueC2 string
}

// HeaderRow documentation
type HeaderRow struct {
	Value string
}

// Value docuemntation
type Value struct {
	//Year  int
	Value float32
}

// BRow documentation
type BRow struct {
	Year    int
	Text    string
	ValueC1 float32
	ValueC2 float32
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
	eusupports := getEUSupports(processid)
	guarantors := getGuarantors(processid)
	maintenancecosts := getMaintenanceCosts(processid)
	//
	processall.Processes = append(processes)
	processall.Applicants = append(applicants)
	processall.Companies = append(companies)
	processall.Collaterals = append(collaterals)
	processall.Loans = append(loans)
	processall.ExtLoans = append(extloans)
	processall.Households = append(households)
	processall.CompanyEconomies = append(companyeconomies)
	processall.PersonalEconomies = append(personaleconomies)
	processall.KycInformations = append(kycinformations)
	processall.Budgets = append(budgets)
	processall.EUSupports = append(eusupports)
	processall.Guarantors = append(guarantors)
	processall.MaintenanceCosts = append(maintenancecosts)
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
	// Create a new documents to receive customer data
	pdf := gofpdf.New("P", "mm", "A4", "")

	hr := HeaderRow{}
	// Processes
	//tr := pdf.UnicodeTranslatorFromDescriptor("")
	cp := pdf.UnicodeTranslatorFromDescriptor("cp1252")
	hr.Value = "Processes"
	tr := []TableRow{} // Initiate table on page one
	for _, process := range processes {
		pdf.AddPage()
		tr = []TableRow{}
		pdf.SetFont("Arial", "B", 16)
		pdf = header1(pdf, hr.Value)
		pdf.SetFont("Arial", "B", 10)
		tr = append(tr, TableRow{Key: "ProcessID:", Value: cp(process.ProcessID)})
		tr = append(tr, TableRow{Key: "Created date:", Value: process.ProcessCreatedDate})
		tr = append(tr, TableRow{Key: "Last Accessed:", Value: process.LastAccessed})
		tr = append(tr, TableRow{Key: "CaseID", Value: process.CaseID})
		tr = append(tr, TableRow{Key: "CaseIDStatus", Value: process.CaseIDStatus})
		tr = append(tr, TableRow{Key: cp("Kunder"), Value: ""})
		for _, customerID := range process.CustomerID {
			tr = append(tr, TableRow{Key: "\tKund:", Value: "\t\t" + customerID.CustomerID})
		}
		pdf = table1(pdf, tr) // add table to page
	}
	// Applicants
	hr = HeaderRow{}
	hr.Value = "Applicants"
	for _, applicant := range applicants {
		tr = []TableRow{}
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf = header1(pdf, hr.Value)
		pdf.SetFont("Arial", "B", 10)
		tr = append(tr, TableRow{Key: "customerId:", Value: applicant.CustomerID})
		tr = append(tr, TableRow{Key: "name:", Value: cp(applicant.ApplicantName)})
		tr = append(tr, TableRow{Key: "Adress:", Value: cp(applicant.ApplicantAddress)})
		tr = append(tr, TableRow{Key: "PostNr:", Value: cp(applicant.ApplicantPostAddress)})
		tr = append(tr, TableRow{Key: cp("Kontaktinformation:*"), Value: ""})
		for _, ci := range applicant.ContactInformation {
			//log.Printf("eMail: %s - MobileNumber: %s", ci.ApplicantEmail, ci.ApplicantMobileNumber)
			tr = append(tr, TableRow{Key: "\teMail:", Value: "\t\t" + cp(ci.ApplicantEmail)})
			tr = append(tr, TableRow{Key: "\tphone:", Value: "\t\t" + ci.ApplicantMobileNumber})
		}
		tr = append(tr, TableRow{Key: "", Value: ""})
		tr = append(tr, TableRow{Key: "Intresenttyp:", Value: cp(applicant.StakeholderType)})
		tr = append(tr, TableRow{Key: "Kontakt via sms:", Value: fmt.Sprintf("%v", applicant.ApplicantBySms)})
		tr = append(tr, TableRow{Key: "Kontakt via eMail:", Value: fmt.Sprintf("%v", applicant.ApplicantByeMail)})
		tr = append(tr, TableRow{Key: cp("Anställd:"), Value: fmt.Sprintf("%v", applicant.ApplicantEmployed)})
		pdf = table1(pdf, tr) // add table to page current page
	}
	//
	// Companies
	hr = HeaderRow{}
	hr.Value = cp("Företag")
	for _, company := range companies {
		pdf.AddPage()
		tr = []TableRow{}
		pdf.SetFont("Arial", "B", 16)
		pdf = header1(pdf, hr.Value)
		pdf.SetFont("Arial", "B", 10)
		tr = append(tr, TableRow{Key: "companyId:", Value: company.CompanyID})
		tr = append(tr, TableRow{Key: "OrgNr:", Value: company.OrgNumber})
		tr = append(tr, TableRow{Key: "CompanyName:", Value: cp(company.CompanyName)})
		//tr = append(tr, TableRow{Key: "BusinessFocus:", Value: cp(company.BusinessFocus)})
		pdf.SetFont("Arial", "B", 16)
		tr = append(tr, TableRow{Key: cp("Affärsinriktning:"), Value: ""})
		pdf.SetFont("Arial", "B", 10)
		for _, bf := range company.BusinessFocuses {
			tr = append(tr, TableRow{Key: "\tID:", Value: "\t" + bf.BusinessID})
			tr = append(tr, TableRow{Key: "\tKategori:", Value: "\t" + cp(bf.BusinessCategory)})
			tr = append(tr, TableRow{Key: "\tInriktning:", Value: "\t" + cp(bf.BusinessDirection)})
			tr = append(tr, TableRow{Key: "\t%-andel:", Value: fmt.Sprintf("\t%v", bf.BusinessPart)})
		}
		tr = append(tr, TableRow{Key: "", Value: ""})
		tr = append(tr, TableRow{Key: "Created:", Value: company.Created})
		tr = append(tr, TableRow{Key: "SNI-code:", Value: company.IndustriCode})
		tr = append(tr, TableRow{Key: "SNI-Text:", Value: cp(company.IndustriText)})
		tr = append(tr, TableRow{Key: "Selected:", Value: fmt.Sprintf("%v", company.SelectedCompany)})
		tr = append(tr, TableRow{Key: cp("Huvudmän"), Value: ""})
		for _, pr := range company.Principals {
			tr = append(tr, TableRow{Key: "\tPersNr/Namn:", Value: fmt.Sprintf("\t%s - %s", pr.CustomerID, pr.PrincipalName)})
			//tr = append(tr, TableRow{Key: "\tNamn:*", Value: "\t" + pr.PrincipalName})
		}
		pdf = table1(pdf, tr) // add table to page current page
	}
	//
	// Collaterals
	//
	hr = HeaderRow{}
	hr.Value = "Collaterals"
	for _, collateral := range collaterals {
		tr = []TableRow{}
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf = header1(pdf, hr.Value)
		pdf.SetFont("Arial", "B", 10)
		tr = append(tr, TableRow{Key: "ProcessID:", Value: collateral.ProcessID})
		tr = append(tr, TableRow{Key: "collateralID:", Value: collateral.CollateralID})
		tr = append(tr, TableRow{Key: "Fastighetskod:", Value: collateral.CollateralCode})
		tr = append(tr, TableRow{Key: "CustomerID:", Value: collateral.CustomerID})
		tr = append(tr, TableRow{Key: "Kommun:", Value: cp(collateral.CollateralMunicipality)})
		tr = append(tr, TableRow{Key: "fastighetsnamn", Value: cp(collateral.CollateralName)})
		tr = append(tr, TableRow{Key: "Gata:", Value: cp(collateral.CollateralStreet)})
		tr = append(tr, TableRow{Key: cp("Använd säkerhet:"), Value: fmt.Sprintf("%v", collateral.UseAsCollateral)})
		tr = append(tr, TableRow{Key: cp("Köpa säkerhet:"), Value: fmt.Sprintf("%v", collateral.BuyCollateral)})
		tr = append(tr, TableRow{Key: cp("Taxeringsägare"), Value: ""})
		for _, taxedOwner := range collateral.TaxedOwners {
			tr = append(tr, TableRow{Key: "\tTaxedID:", Value: "\t" + taxedOwner.TaxedID})
			tr = append(tr, TableRow{Key: cp("\tÄgare:"), Value: "\t" + cp(taxedOwner.TaxedOwner)})
		}
		pdf = table1(pdf, tr) // add table to page current page
	}
	//
	// Loans
	//
	hr = HeaderRow{}
	hr.Value = "Loans"
	for _, loan := range loans {
		tr = []TableRow{}
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf = header1(pdf, hr.Value)
		pdf.SetFont("Arial", "B", 11)
		tr = append(tr, TableRow{Key: cp("Lånenummer:"), Value: loan.LoanNumber})
		tr = append(tr, TableRow{Key: cp("Ändamål:"), Value: cp(loan.PurposeOfLoan)})
		tr = append(tr, TableRow{Key: cp("Lånebelopp:"), Value: fmt.Sprintf("%.2f", loan.LoanAmount)})
		tr = append(tr, TableRow{Key: cp("Syfte respektive lån"), Value: ""})
		for _, aim := range loan.Aims {
			tr = append(tr, TableRow{Key: "\tSyfte:", Value: "\t" + cp(aim.AimText)})
			tr = append(tr, TableRow{Key: cp("\tLånebelopp:"), Value: fmt.Sprintf("\t%.2f", aim.LoanAmountPart)})
		}
		pdf = table1(pdf, tr) // add table to page current page
	}
	//
	// ExtLoans
	//
	hr = HeaderRow{}
	hr.Value = "ExtLoans"
	for _, extloan := range extloans {
		tr = []TableRow{}
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf = header1(pdf, hr.Value)
		pdf.SetFont("Arial", "B", 11)
		tr = append(tr, TableRow{Key: cp("ExtLoanId:"), Value: cp(extloan.ExtLoanID)})
		tr = append(tr, TableRow{Key: cp("Institut:"), Value: cp(extloan.ExtCreditInstitut)})
		tr = append(tr, TableRow{Key: cp("Clearing:"), Value: extloan.ExtLoanClearing})
		tr = append(tr, TableRow{Key: cp("Lånenummer:"), Value: extloan.ExtLoanNumber})
		tr = append(tr, TableRow{Key: cp("Belopp:"), Value: fmt.Sprintf("%.2f", extloan.ExtLoanAmount)})
		tr = append(tr, TableRow{Key: cp("Lösa lån:"), Value: fmt.Sprintf("%v", extloan.ExtRedeemLoan)})
		tr = append(tr, TableRow{Key: cp("Lånet står på"), Value: ""})
		for _, owner := range extloan.ExtLoanOwners {
			tr = append(tr, TableRow{Key: cp("\tPersonnummer:"), Value: "\t" + owner.CustomerID})
		}
		pdf = table1(pdf, tr) // add table to page current page
	}
	//
	// EU-Support
	//
	hr = HeaderRow{}
	hr.Value = cp("EU-Stöd")
	//
	for _, eusupport := range eusupports {
		tr = []TableRow{}
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf = header1(pdf, hr.Value)
		pdf.SetFont("Arial", "B", 11)
		tr = append(tr, TableRow{Key: "ProcessID:", Value: eusupport.ProcessID})
		tr = append(tr, TableRow{Key: "EUID:", Value: eusupport.EUID})
		tr = append(tr, TableRow{Key: "EU-Typ:", Value: cp(eusupport.EUType)})
		tr = append(tr, TableRow{Key: cp("År:"), Value: eusupport.SupportYear})
		tr = append(tr, TableRow{Key: cp("Stöd:"), Value: fmt.Sprintf("%.2f", eusupport.SupportAmount)})
		pdf = table1(pdf, tr) // add table to page current page
	}
	//
	// Borgensmän
	//
	hr = HeaderRow{}
	hr.Value = "Borgensman"
	//
	for _, guarantor := range guarantors {
		tr = []TableRow{}
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf = header1(pdf, hr.Value)
		pdf.SetFont("Arial", "B", 11)
		tr = append(tr, TableRow{Key: "ProcessID:", Value: guarantor.ProcessID})
		tr = append(tr, TableRow{Key: "BorgensmanId:", Value: guarantor.GuarantorID})
		tr = append(tr, TableRow{Key: "CustomerId:", Value: guarantor.GuarantorCustomerID})
		tr = append(tr, TableRow{Key: "Namn:", Value: cp(guarantor.GuarantorName)})
		tr = append(tr, TableRow{Key: "Phone:", Value: guarantor.GuarantorPhone})

		pdf = table1(pdf, tr) // add table to page current page
	}
	//
	// MainteanceCosts
	//
	hr = HeaderRow{}
	hr.Value = "Andra boenden"
	//
	for _, maintenancecost := range maintenancecosts {
		tr = []TableRow{}
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf = header1(pdf, hr.Value)
		pdf.SetFont("Arial", "B", 11)
		tr = append(tr, TableRow{Key: "ProcessID:", Value: maintenancecost.ProcessID})
		tr = append(tr, TableRow{Key: "BorgensamnId:", Value: maintenancecost.MaintenanceCostID})
		//
		tr = append(tr, TableRow{Key: cp("Hushåll"), Value: ""})
		for _, typeofhouse := range maintenancecost.TypeOfHouses {
			tr = append(tr, TableRow{Key: "\tHouseId:", Value: "\t" + typeofhouse.HouseID})
			tr = append(tr, TableRow{Key: cp("\ttype av hushåll:"), Value: "\t" + cp(typeofhouse.TypeOfHouse)})
			tr = append(tr, TableRow{Key: cp("\tLån i annan bank:"), Value: fmt.Sprintf("\t%v", typeofhouse.LoanInOtherInstitut)})
			tr = append(tr, TableRow{Key: cp("\tLösa lån:"), Value: fmt.Sprintf("\t%v", typeofhouse.RedeemLoan)})
			tr = append(tr, TableRow{Key: "\tKreditinstitut:", Value: "\t" + cp(typeofhouse.CreditInstitut)})
			tr = append(tr, TableRow{Key: "\tClearing:", Value: "\t" + cp(typeofhouse.LoanClearing)})
			tr = append(tr, TableRow{Key: cp("\tLånenummer:"), Value: "\t" + cp(typeofhouse.InstitutLoanNumber)})
			tr = append(tr, TableRow{Key: cp("\tLånebelopp:"), Value: "\t" + fmt.Sprintf("%.2f", typeofhouse.LoanAmount)})
			//tr = append(tr, TableRow{Key: cp("Ägare:"), Value: cp(typeofhouse.LoanOwner)})
			tr = append(tr, TableRow{Key: cp("Ägare"), Value: ""})
			for _, loanOwner := range typeofhouse.LoanOwners {
				tr = append(tr, TableRow{Key: "\t\tKund:", Value: "\t\t" + loanOwner.CustomerID})
			}
			tr = append(tr, TableRow{Key: "", Value: ""})
			//tr = append(tr, TableRow{Key: cp("\tStöd:"), Value: fmt.Sprintf("\t%.2f", typeofhouse.LoanAmount)})
			tr = append(tr, TableRow{Key: cp("\tDriftkostnad:"), Value: fmt.Sprintf("\t%.2f", typeofhouse.MaintenanceCost)})

		}
		pdf = table1(pdf, tr) // add table to page current page
	}
	//
	// Budget
	//
	y1 := 0
	y2 := 0
	i := 0
	//
	hr = HeaderRow{}
	hr.Value = "Budget"
	trb := []BudgetTableRow{} // Initiate table on page one
	brows := []BRow{}
	for _, budget := range budgets {
		if i == 0 {
			pdf.AddPage()
			pdf.SetFont("Arial", "B", 16)
			pdf = header1(pdf, hr.Value)
			pdf.SetFont("Arial", "B", 11)
			brows = fixBudgetRows(budget.BudgetYears)
			//y1 = budget.BudgetYear
		} else {
			//y2 = budget.BudgetYears
			//brows = fixBudgetRows(budget.BudgetYears)
		}
		trb = append(trb, BudgetTableRow{Key: cp("CompanyID:"), Text: budget.CompanyEconomyID})
		i++
	}
	//
	j := 0
	trb = append(trb, BudgetTableRow{
		//Key:     fmt.Sprintf("%d", i),
		//Text:    fmt.Sprintf("%s", br.Text),
		ValueC1: fmt.Sprintf("%v", y1),
		ValueC2: fmt.Sprintf("%v", y2)})
	//
	for _, br := range brows {
		//fmt.Println(fmt.Sprintf("%d: - Text: %s C1: %v C2: %v", j+1, br.Text, br.ValueC1, br.ValueC2))
		trb = append(trb, BudgetTableRow{
			Key:     fmt.Sprintf("%d", j+1),
			Text:    fmt.Sprintf("%s", cp(br.Text)),
			ValueC1: fmt.Sprintf("%v", br.ValueC1),
			ValueC2: fmt.Sprintf("%v", br.ValueC2)})
		j++
	}
	//
	pdf = tablebudget(pdf, trb) // add table to page current page
	//}
	//pdf = table1(pdf, tr) // add table to page current page
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

// GetEUSupports documentation
func getEUSupports(processid string) []models.EUSupportType {

	//message := models.MessageBody{}
	eusupports := []models.EUSupportType{}
	//euret := make([]models.EUSupportType, 1, 1)
	var euret []models.EUSupportType
	//
	file, err := ioutil.ReadFile("json/eusupports.json")
	if err != nil {
		fmt.Println("Error reading eusupports.json - %s", err)
	}
	_ = json.Unmarshal([]byte(file), &eusupports)
	//
	for _, eusupport := range eusupports {
		if eusupport.ProcessID == processid {
			//fmt.Println("Id: %s", fmt.Sprintf("%v", eusupport.SupportAmount))
			euret = append(euret, eusupport)
		}
	}
	return euret
}

// getGuarantors documentation
func getGuarantors(processid string) []models.GuarantorType {

	//message := models.MessageBody{}
	guarantors := []models.GuarantorType{}
	//budret := make([]models.BudgetType, 2, 2)
	var guaret []models.GuarantorType
	//
	file, err := ioutil.ReadFile("json/guarantors.json")
	if err != nil {
		fmt.Println("Error reading guarantors.json - %s", err)
	}
	_ = json.Unmarshal([]byte(file), &guarantors)
	//
	for _, guarantor := range guarantors {
		if guarantor.ProcessID == processid {
			//budret[j] = budget
			guaret = append(guaret, guarantor)
		}
	}
	return guaret
}

// getMaintenanceCosts documentation
func getMaintenanceCosts(processid string) []models.MaintenanceCostType {

	//message := models.MessageBody{}
	maintenancecosts := []models.MaintenanceCostType{}
	//budret := make([]models.BudgetType, 2, 2)
	var mainret []models.MaintenanceCostType
	//
	file, err := ioutil.ReadFile("json/maintenancecosts.json")
	if err != nil {
		fmt.Println("Error reading maintenancecosts.json - %s", err)
	}
	_ = json.Unmarshal([]byte(file), &maintenancecosts)
	//
	for _, maintenancecost := range maintenancecosts {
		if maintenancecost.ProcessID == processid {
			//budret[j] = budget
			mainret = append(mainret, maintenancecost)
		}
	}
	return mainret
}

// savePDF documentation
func savePDF(filename string, pdf *gofpdf.Fpdf) error {
	return pdf.OutputFileAndClose(filename)
}

// header documentation
func header1(pdf *gofpdf.Fpdf, hdr string) *gofpdf.Fpdf {
	pdf.SetFont("Arial", "B", 12)
	pdf.SetFillColor(240, 240, 240)
	pdf.CellFormat(185, 7, hdr, "1", 0, "", true, 0, "")
	pdf.Ln(-1)
	return pdf
}

// table1 documentation
func table1(pdf *gofpdf.Fpdf, tbl []TableRow) *gofpdf.Fpdf {
	// Reset font and fill color.
	pdf.SetFont("Arial", "", 12)
	pdf.SetFillColor(255, 255, 255)
	// Every column gets aligned according to its contents.
	//align := []string{"L", "C", "L", "R", "R", "R"}
	for _, line := range tbl {
		pdf.CellFormat(60, 10, line.Key, "1", 0, "L", false, 0, "")
		pdf.CellFormat(125, 10, line.Value, "1", 0, "L", false, 0, "")

		pdf.Ln(-1)
	}
	return pdf
}

// tablebudget documentation
func tablebudget(pdf *gofpdf.Fpdf, tbl []BudgetTableRow) *gofpdf.Fpdf {
	// Reset font and fill color.
	pdf.SetFont("Arial", "", 12)
	pdf.SetFillColor(255, 255, 255)
	// Every column gets aligned according to its contents.
	//align := []string{"L", "C", "L", "R", "R", "R"}
	for _, line := range tbl {
		pdf.CellFormat(25, 10, line.Key, "1", 0, "L", false, 0, "")
		pdf.CellFormat(120, 10, line.Text, "1", 0, "L", false, 0, "")
		pdf.CellFormat(20, 10, line.ValueC1, "1", 0, "R", false, 0, "")
		pdf.CellFormat(20, 10, line.ValueC2, "1", 0, "R", false, 0, "")
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

// fixBudgetRows documentation
func fixBudgetRows(budgets []models.BudgetYear) []BRow {

	// BRow documentation
	fmt.Println(fmt.Sprintf("%v", getColumnText(24)))
	//BRows := []BRow{}
	//j := 0
	//BRows = make([]BRow, 0, 25)
	var BRows []BRow
	var brow BRow
	//Flytta ut år 1 samt år 2 till variabler och hämta år 1 samt år två i loopen 0 - 24
	// Första loopen ska alltså försvinna
	y1 := 0
	y2 := 0
	i := 0
	for _, budget := range budgets {
		if i == 0 {
			y1 = budget.BudgetYear
		} else {
			y2 = budget.BudgetYear
		}
		i++
	}
	//fmt.Println(fmt.Sprintf("Len:%v", len(budgets)))
	for j := 0; j <= 24; j++ {
		if j == 0 {
			//brow.Year = y1
		}
		brow.Year = y1
		brow.Text = getColumnText(j)
		brow.ValueC1 = getColumn1Value(j, y1, budgets)
		brow.ValueC2 = getColumn2Value(j, y2, budgets)
		BRows = append(BRows, brow)
	}
	//
	return BRows
}

// getColumn1Value documentation
func getColumn1Value(index int, year int, budgets []models.BudgetYear) float32 {

	var retval float32
	for _, budget := range budgets {
		if budget.BudgetYear == year {
			for _, val := range budget.Values {
				switch index {
				case 0:
					retval = val.Value1
				case 1:
					retval = val.Value2
				case 2:
					retval = val.Value3
				case 3:
					retval = val.Value4
				case 4:
					retval = val.Value5
				case 5:
					retval = val.Value6
				case 6:
					retval = val.Value7
				case 7:
					retval = val.Value8
				case 8:
					retval = val.Value9
				case 9:
					retval = val.Value10
				case 10:
					retval = val.Value11
				case 11:
					retval = val.Value12
				case 12:
					retval = val.Value13
				case 13:
					retval = val.Value14
				case 14:
					retval = val.Value15
				case 15:
					retval = val.Value16
				case 16:
					retval = val.Value17
				case 17:
					retval = val.Value18
				case 18:
					retval = val.Value19
				case 19:
					retval = val.Value20
				case 20:
					retval = val.Value21
				case 21:
					retval = val.Value22
				case 22:
					retval = val.Value23
				case 23:
					retval = val.Value24
				case 24:
					retval = val.Value25
				} //i++
			}
		}
	}
	return retval
}

// getColumn2Value documentation
func getColumn2Value(index int, year int, budgets []models.BudgetYear) float32 {

	var retval float32
	//i := 0
	for _, budget := range budgets {
		if budget.BudgetYear == year {
			//i = 0
			for _, val := range budget.Values {
				switch index {
				case 0:
					retval = val.Value1
				case 1:
					retval = val.Value2
				case 2:
					retval = val.Value3
				case 3:
					retval = val.Value4
				case 4:
					retval = val.Value5
				case 5:
					retval = val.Value6
				case 6:
					retval = val.Value7
				case 7:
					retval = val.Value8
				case 8:
					retval = val.Value9
				case 9:
					retval = val.Value10
				case 10:
					retval = val.Value11
				case 11:
					retval = val.Value12
				case 12:
					retval = val.Value13
				case 13:
					retval = val.Value14
				case 14:
					retval = val.Value15
				case 15:
					retval = val.Value16
				case 16:
					retval = val.Value17
				case 17:
					retval = val.Value18
				case 18:
					retval = val.Value19
				case 19:
					retval = val.Value20
				case 20:
					retval = val.Value21
				case 21:
					retval = val.Value22
				case 22:
					retval = val.Value23
				case 23:
					retval = val.Value24
				case 24:
					retval = val.Value25
				}
				//i++
			}
		}
	}
	return retval
}

// getColumnText documentation
func getColumnText(index int) string {

	text := []models.TextValue{
		{"Skog"},
		{"Växtodling"},
		{"EU-stöd"},
		{"Övrig djurproduktion"},
		{"Förändring av lager produktion"},
		{"Mjölk"},
		{"Övriga intäkter"},
		{"Omsättning totalt"},
		{"Inköp (Råvaror och förnödenheter)"},
		{"Arrendekostnader"},
		{"Personalkostnader"},
		{"Övriga rörelsekostnader"},
		{"S:a kostnader (summa value9-12)"},
		{"Resultat före avskrivningar (value8 + (-value13))"},
		{"Avskrivning inventarier (exkl. byggnadsinventerier)"},
		{"Övriga avskrivningar"},
		{"S:a avskrivningar (-value15) + (-value16)"},
		{"Resultat före avskrivningar value14 + (-value17)"},
		{"Finansiella intäkter"},
		{"Finansiella kostnader"},
		{"Resultat finansiella poster (value18) + value19 + (-value20)"},
		{"Extraordinärar intäkter och kostnader"},
		{"Bokslutsdispositioner"},
		{"Skatt (ägaruttag prognosår EF)"},
		{"Åretsresultat (value25) sum (value22+value23+value24)"},
	}
	return text[index].Text
}
