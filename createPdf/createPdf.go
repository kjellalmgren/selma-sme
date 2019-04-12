package createpdf

import (
	"fmt"
	"log"
	"selmasme/models"
	"time"

	"github.com/jung-kurt/gofpdf"
)

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
	pdf.Cell(40, 10, "Applicant")
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(50, 10, applicant.CustomerID)
	pdf.Cell(60, 10, applicant.ApplicantName)
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
