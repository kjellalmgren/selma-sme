package submitapplications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	createpdf "selmasme/createPdf"
	"selmasme/models"
)

// Submit documentation
func Submit(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	processid := r.Header.Get("X-process-Id")
	var data models.SubmitApplication
	//sfile := "sme.pdf"
	//
	var r1 []byte
	r1, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Printf("Submit Applications for %s\r\n", processid)
	json.NewDecoder(bytes.NewReader([]byte(r1))).Decode(&data)
	fmt.Printf("\r\nr1: %s\r", r1)
	//fmt.Printf("\r\nsubmit: %v\r", data.SubmitApplication)
	//
	if data.SubmitApplication == false {
		err1 := models.Error{}
		err1.Status = http.StatusNotFound
		err1.Message = "Application is uncomplete..."
		err1.Error = "APPLICATION_NOT_COMPLETE"
		if err := json.NewEncoder(w).Encode(err1); err != nil {
			w.WriteHeader(http.StatusNotFound)
			//panic(err)
		}
		w.WriteHeader(http.StatusNotFound)
	} else if data.SubmitApplication == true {
		//
		mb := createpdf.CreatePdfDocument(processid)
		if mb.Status == true {
			err1 := models.Error{}
			err1.Status = http.StatusOK
			err1.Message = fmt.Sprintf("Application is transfered...%s", mb.Filename)
			err1.Error = "APPLICATION_CREATED"
			if err := json.NewEncoder(w).Encode(err1); err != nil {
				w.WriteHeader(http.StatusNotFound)
				//panic(err)
			}
		}
		//fmt.Fprintf(w, "pdf document has been created - %s", sfile)
		w.WriteHeader(http.StatusOK)
		//
	}
}
