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

// CreatePdfDocuments documentation
func CreatePdfDocument(processid string) models.MessageBody {

	message := models.MessageBody{}

	//processall := models.ProcessAllType{}
	// processes := []models.ProcessType{}
	//applicants := []models.ApplicantType{}
	// loans := []models.LoanType{}
	// extloans := []models.ExtLoanType{}
	// households := []models.HouseholdType{}
	// companies := []models.CompanyType{}
	// companyeconomies := []models.CompanyEconomyType{}
	// personaleconomies := []models.PersonalEconomyType{}
	// kycinformations := []models.KycInformationType{}
	// collaterals := []models.CollateralType{}
	// budgets := []models.BudgetType{}

	processall := models.ProcessAllType{}
	processes := getProcesses(processid)
	applicants := getApplicants(processid)

	processall.Processes = append(processes)
	processall.Applicans = append(applicants)
	//

	return message

}

//
// getProcesses
func getProcesses(processid string) []models.ProcessType {

	processes := []models.ProcessType{}
	procret := make([]models.ProcessType, 1, 2)

	file, err := ioutil.ReadFile("json/processes.json")
	if err != nil {
		fmt.Println("Error reading processes.json - %s", err)
	}
	_ = json.Unmarshal([]byte(file), &processes)
	j := 0
	for i := 0; i < len(processes); i++ {
		if processes[i].ProcessID == processid {
			procret[j] = processes[i]
			j++
		}
	}
	return processes
}

//
// getApplicant documentation
func getApplicants(processid string) []models.ApplicantType {

	//message := models.MessageBody{}
	applicants := []models.ApplicantType{}
	appret := make([]models.ApplicantType, 1, 5)
	//
	file, err := ioutil.ReadFile("json/applicants.json")
	if err != nil {
		fmt.Println("Error reading applicants.json - %s", err)
	}
	_ = json.Unmarshal([]byte(file), &applicants)
	//
	j := 0
	for i := 0; i < len(applicants); i++ {
		if applicants[i].ProcessID == processid {
			appret[j] = applicants[i]
			j++
		}
	}
	return applicants
}

// Should be replace by CreatePdfDocument
// CreatePdf documentation
func CreatePdf(applicant models.ApplicantType) models.MessageBody {

	message := models.MessageBody{}

	t := time.Now()
	fmt.Println(t.String())
	fmt.Println(t.Format("2006-01-02 15:04:05"))

	filename := applicant.CustomerID + "-" + t.Format("2006-01-02T15:04:05"+".pdf")
	fmt.Println(filename)
	//filename = "sme.pdf"
	//
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	//str := []string{}
	strHeader := make([]string, 1, 1)
	strHeader[0] = "Applicants"
	strTable := make([][]string, 1, 20)
	pdf = header(pdf, strHeader)
	pdf.SetFont("Arial", "B", 12)

	var tr []TableRow

	tr = append(tr, TableRow{Key: "customerId:", Value: applicant.CustomerID})
	tr = append(tr, TableRow{Key: "name:", Value: applicant.ApplicantName})
	tr = append(tr, TableRow{Key: "Adress:", Value: applicant.ApplicantAddress})
	tr = append(tr, TableRow{Key: "PostNr:", Value: applicant.ApplicantPostAddress})
	for _, ci := range applicant.ContactInformation {
		log.Printf("eMail: %s - MobileNumber: %s", ci.ApplicantEmail, ci.ApplicantMobileNumber)
		tr = append(tr, TableRow{Key: "eMail:", Value: ci.ApplicantEmail})
		tr = append(tr, TableRow{Key: "eMail:", Value: ci.ApplicantMobileNumber})
	}
	tr = append(tr, TableRow{Key: "Stakeholder:", Value: applicant.StakeholderType})
	tr = append(tr, TableRow{Key: "Kontakt via sms:", Value: fmt.Sprintf("%v", applicant.ApplicantBySms)})
	tr = append(tr, TableRow{Key: "Kontakt via eMail:", Value: fmt.Sprintf("%v", applicant.ApplicantByeMail)})
	tr = append(tr, TableRow{Key: "Anställd:", Value: fmt.Sprintf("%v", applicant.ApplicantEmployeed)})

	//str[5] = applicant.ApplicantByeMail.string()
	//str[6] = applicant.ApplicantBySms.string()
	pdf = table(pdf, strTable[1:])
	pdf = table1(pdf, tr)
	//pdf.Cell(50, 10, applicant.CustomerID)
	//pdf.Cell(60, 10, applicant.ApplicantName)
	//
	err := savePDF(filename, pdf)
	if err != nil {
		log.Fatalf("createPdf::CreatePdf: Cannot save PDF: %s|n", filename)
		message.Status = false
		message.MessageText = "Error creating pdf document..."
		message.Filename = ""
	} else {
		message.Status = true
		message.MessageText = "OK"
		message.Filename = ""
	}
	//
	return message
}

// savePDF documentation
func savePDF(filename string, pdf *gofpdf.Fpdf) error {
	return pdf.OutputFileAndClose(filename)
}

// header documentation
func header(pdf *gofpdf.Fpdf, hdr []string) *gofpdf.Fpdf {
	pdf.SetFont("Arial", "B", 16)
	pdf.SetFillColor(240, 240, 240)
	for _, str := range hdr {
		pdf.CellFormat(40, 7, str, "1", 0, "", true, 0, "")
	}
	pdf.Ln(-1)
	return pdf
}

// table documentation
func table(pdf *gofpdf.Fpdf, tbl [][]string) *gofpdf.Fpdf {
	// Reset font and fill color.
	pdf.SetFont("Arial", "", 16)
	pdf.SetFillColor(255, 255, 255)
	// Every column gets aligned according to its contents.
	align := []string{"L", "C", "L", "R", "R", "R"}
	for _, line := range tbl {
		for i, str := range line {
			pdf.CellFormat(40, 7, str, "1", 0, align[i], false, 0, "")
		}
		pdf.Ln(-1)
	}
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
