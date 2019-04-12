package createpdf

import (
	"fmt"
	"log"
	"selmasme/models"
	"time"

	"github.com/jung-kurt/gofpdf"
)

// CreatePdf documentation
func CreatePdf(customerID string) models.MessageBody {

	message := models.MessageBody{}

	t := time.Now()
	fmt.Println(t.String())
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	filename := customerID + "-" + t.Format("2006-01-02T15:04:05")
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Applicants-1")
	//
	err := savePDF(filename, pdf)
	if err != nil {
		log.Fatalf("CreatePdf::CreatePdf: Cannot save PDF: %s|n", filename)
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
